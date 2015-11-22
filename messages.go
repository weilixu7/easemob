package easemob

import (
	"net/http"
)

/**
 * MessageService handles communication with the message related methods of
 * the Easemo API
 *
 * Easemob API docs: http://docs.easemob.com/doku.php?id=start:100serverintegration:50messages
 */
type MessagesService struct {
	client *Client
}

/**
 * Send text message to users
 *
 * http://docs.easemob.com/doku.php?id=start:100serverintegration:50messages#发送文本消息
 */
func (s *MessagesService) SendTextMessagesToUsers(from string, text string,
	userIds ...string) (*Response, error) {

	var (
		err        error
		putOptions *MessagePutOptions
		path       string
		req        *http.Request
		resp       *Response
	)

	putOptions = &MessagePutOptions{
		TargetType: "users",
		Target:     userIds,
		Msg:        &MessageType{Type: "txt", Msg: text},
		From:       from,
	}

	path = "messages"

	req, err = s.client.NewRequest("POST", path, putOptions)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.Do(req)
	return resp, err
}
