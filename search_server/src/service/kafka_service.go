package service

import "line_china/search_server/src/provider"

// SendMsgKafka 向kafka发送消息
func SendMsgKafka(publishId int, msg []byte) error {
	return provider.AsyncProducer(publishId, msg)
}
