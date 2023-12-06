package goshopee

type PaymentService interface {
	GetEscrowDetail(uint64, string, string) (*GetEscrowDetailResponse, error)
}

type PaymentServiceOp struct {
	client *Client
}

type GetEscrowDetailRequest struct {
	OrderSN string `url:"order_sn"`
}

type GetEscrowDetailResponse struct {
	BaseResponse

	Response GetEscrowDetailResponseData `json:"response"`
}

type GetEscrowDetailResponseData struct {
	OrderSn           string   `json:"order_sn"`
	BuyerUserName     string   `json:"buyer_user_name"`
	ReturnOrderSnList []string `json:"return_order_sn_list"`
}

func (s *PaymentServiceOp) GetEscrowDetail(sid uint64, orderSn string, tok string) (*GetEscrowDetailResponse, error) {
	path := "/payment/get_escrow_detail"

	opt := GetEscrowDetailRequest{
		OrderSN: orderSn,
	}

	resp := new(GetEscrowDetailResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}
