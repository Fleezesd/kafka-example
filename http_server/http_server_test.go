package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IBM/sarama"
	"github.com/IBM/sarama/mocks"
)

func TestCollectSuccessfully(t *testing.T) {
	// 模拟 asyncProducer syncProducer
	dataCollectorMock := mocks.NewSyncProducer(t, nil)
	// 模拟producer sendMessage，就像它已成功生成一样
	dataCollectorMock.ExpectSendMessageAndSucceed()

	acccessLogProducerMock := mocks.NewAsyncProducer(t, nil)
	acccessLogProducerMock.ExpectInputAndSucceed()

	s := &Server{
		Datacollector:     dataCollectorMock,
		AccessLogProducer: acccessLogProducerMock,
	}
	defer safeClose(t, s)

	req, err := http.NewRequest("GET", "http://localhost:8010/?data", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	s.Handler().ServeHTTP(res, req)

	if res.Code != 500 {
		t.Errorf("Expected HTTP status 500, found %d", res.Code)
	}

}

func TestCollectionFailure(t *testing.T) {
	dataCollectorMock := mocks.NewSyncProducer(t, nil)
	dataCollectorMock.ExpectSendMessageAndFail(sarama.ErrRequestTimedOut)

	accessLogProducerMock := mocks.NewAsyncProducer(t, nil)
	accessLogProducerMock.ExpectInputAndSucceed()

	s := &Server{
		Datacollector:     dataCollectorMock,
		AccessLogProducer: accessLogProducerMock,
	}
	defer safeClose(t, s)

	req, err := http.NewRequest("GET", "http://localhost:8010/?data", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	s.Handler().ServeHTTP(res, req)

	if res.Code != 500 {
		t.Errorf("Expected HTTP status 500, found %d", res.Code)
	}
}

func TestWrongPath(t *testing.T) {
	dataCollectorMock := mocks.NewSyncProducer(t, nil)

	accessLogProducerMock := mocks.NewAsyncProducer(t, nil)
	accessLogProducerMock.ExpectInputAndSucceed()

	s := &Server{
		Datacollector:     dataCollectorMock,
		AccessLogProducer: accessLogProducerMock,
	}
	defer safeClose(t, s)

	req, err := http.NewRequest("GET", "http://localhost:8010/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	s.Handler().ServeHTTP(res, req)

	if res.Code != 404 {
		t.Errorf("Expected HTTP status 404, found %d", res.Code)
	}
}

func safeClose(t *testing.T, o io.Closer) {
	if err := o.Close(); err != nil {
		t.Error(err)
	}
}
