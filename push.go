package goshopee

type PushService interface {
	GetLostPushMessage() (*GetLostPushMessageResponse, error)
	ConfirmConsumedLostPushMessage(lastMessageId uint64) (*BaseResponse, error)
}

type PushServiceOp struct {
	client *Client
}

type GetLostPushMessageResponse struct {
	BaseResponse

	Response GetLostPushMessage `json:"response"`
}

type GetLostPushMessage struct {
	HasNextPage   bool   `json:"has_next_page"`
	LastMessageId uint64 `json:"last_message_id"`

	PushMessageList []PushMessage `json:"push_message_list"`
}

type PushMessage struct {
	ShopId    uint64 `json:"shop_id"`
	Code      int    `json:"code"`
	Timestamp uint64 `json:"timestamp"`
	Data      string `json:"data"`
}

func (s *PushServiceOp) GetLostPushMessage() (*GetLostPushMessageResponse, error) {
	path := "/push/get_lost_push_message"

	resp := new(GetLostPushMessageResponse)
	err := s.client.Get(path, resp, nil)
	return resp, err
}

func (s *PushServiceOp) ConfirmConsumedLostPushMessage(lastMessageId uint64) (*BaseResponse, error) {
	path := "/push/confirm_consumed_lost_push_message"

	resp := new(BaseResponse)
	req := map[string]interface{}{
		"last_message_id": lastMessageId,
	}

	err := s.client.Post(path, req, resp)
	return resp, err
}
