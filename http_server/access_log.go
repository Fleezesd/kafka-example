package main

import "encoding/json"

// accesslog 写入kafka 预定义结构
type accessLogEntry struct {
	Method       string  `json:"method"`
	Host         string  `json:"host"`
	Path         string  `json:"path"`
	IP           string  `json:"ip"`
	ResponseTime float64 `json:"response_time"`

	encoded []byte
	err     error
}

//		value ----> type Encoder interface {
//			Encode() ([]byte, error)
//			Length() int
//		}
//	 Encoder一个适用于任何类型的简单接口，可以将其编码为字节数组，以便作为 Kafka 消息的键或值发送。
//	 Length() 是作为优化而提供的，
//	 并且必须在 Encode() 的结果上返回与 len() 相同的值
func (ale *accessLogEntry) Length() int {
	if ale.encoded == nil && ale.err == nil {
		ale.encoded, ale.err = json.Marshal(ale)
	}
	return len(ale.encoded)
}

func (ale *accessLogEntry) Encode() ([]byte, error) {
	if ale.encoded == nil && ale.err == nil {
		ale.encoded, ale.err = json.Marshal(ale)
	}
	return ale.encoded, ale.err
}
