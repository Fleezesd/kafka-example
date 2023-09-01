package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IBM/sarama"
)

type Server struct {
	Datacollector     sarama.SyncProducer  // 同步producer
	AccessLogProducer sarama.AsyncProducer // 异步producer
}

// 启动kafka server http启动方式
func (s *Server) Run(addr string) error {
	httpServer := &http.Server{
		Addr:    addr,
		Handler: s.Handler(),
	}

	log.Printf("Listening for requests on %s...\n", addr)
	return httpServer.ListenAndServe()
}

func (s *Server) Handler() http.Handler {
	return s.withAccessLog(s.CollectStringData())
}

// 访问日志信息 不在乎是否同步发送到 Kafka asyncProducer
func (s *Server) withAccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()

		// sync producer httpserver start
		next.ServeHTTP(w, r)

		entry := &accessLogEntry{
			Method:       r.Method,
			Host:         r.Host,
			Path:         r.RequestURI,
			IP:           r.RemoteAddr,
			ResponseTime: float64(time.Since(started)) / float64(time.Second),
		}

		// 构建 key-value 结构 想存入kafka的信息
		s.AccessLogProducer.Input() <- &sarama.ProducerMessage{
			Topic: "access_log",
			Key:   sarama.StringEncoder(r.RemoteAddr),
			Value: entry,
		}
	})
}

// http 响应 来看消息是否成功发送到 Kafka syncProducer
func (s *Server) CollectStringData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test" {
			http.NotFound(w, r)
			return
		}

		// 因为是syncProducer 即为三个返回值 若为async则为channel形式返回
		partition, offset, err := s.Datacollector.SendMessage(&sarama.ProducerMessage{
			Topic: "important",
			Value: sarama.StringEncoder(r.URL.RawQuery),
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to store your data: %s", err)
		} else {
			// 分区和偏移量信息
			fmt.Fprintf(w, "Your data is stored with unique identifier important/%d/%d", partition, offset)
		}
	})
}

// 关闭各producer
func (s Server) Close() error {
	if err := s.Datacollector.Close(); err != nil {
		log.Println("Failed to shut down data collector cleanly", err)
		return err
	}
	if err := s.AccessLogProducer.Close(); err != nil {
		log.Println("Failed to shut down access log producer cleanly", err)
		return err
	}
	return nil
}
