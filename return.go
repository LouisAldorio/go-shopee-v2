package goshopee

type ReturnService interface {
	GetReturnList(sid uint64, item GetReturnListRequest, tok string) (*GetReturnListResponse, error)
}

type ReturnServiceOp struct {
	client *Client
}

type GetReturnListRequest struct {
	Page     int     `url:"page_no"`
	PageSize int     `url:"page_size"`
	Status   *string `url:"status,omitempty"`

	Start *uint64 `url:"create_time_from,omitempty"`
	End   *uint64 `url:"create_time_to,omitempty"`
}

type GetReturnListResponseData struct {
	More bool `json:"more"`

	Returns []Return `json:"return"`
}

type Return struct {
	ReturnSN                 string       `json:"return_sn"`
	Image                    []string     `json:"image"`
	Reason                   string       `json:"reason"`
	TextReason               string       `json:"text_reason"`
	RefundAmount             float64      `json:"refund_amount"`
	Currency                 string       `json:"currency"`
	CreateTime               uint64       `json:"create_time"`
	UpdateTime               uint64       `json:"update_time"`
	Status                   string       `json:"status"`
	DueDate                  uint64       `json:"due_date"`
	TrackingNumber           string       `json:"tracking_number"`
	DisputeReason            []string     `json:"dispute_reason"`
	DisputeTextReason        []string     `json:"dispute_text_reason"`
	NeedsLogistics           bool         `json:"needs_logistics"`
	AmountBeforeDiscount     float64      `json:"amount_before_discount"`
	User                     ReturnUser   `json:"user"`
	Items                    []ReturnItem `json:"item"`
	OrderSN                  string       `json:"order_sn"`
	ReturnShipDueDate        uint64       `json:"return_ship_due_date"`
	ReturnSellerDueDate      uint64       `json:"return_seller_due_date"`
	NegotiationStatus        string       `json:"negotiation_status"`
	SellerProofStatus        string       `json:"seller_proof_status"`
	SellerCompensationStatus string       `json:"seller_compensation_status"`
}

type ReturnUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Portrait string `json:"portrait"`
}

type ReturnItem struct {
	ModelId      uint64   `json:"model_id"`
	Name         string   `json:"name"`
	Images       []string `json:"images"`
	Amount       int      `json:"amount"`
	ItemPrice    float64  `json:"item_price"`
	IsAddOnDeal  bool     `json:"is_add_on_deal"`
	IsMainItem   bool     `json:"is_main_item"`
	AddOnDealId  uint64   `json:"add_on_deal_id"`
	ItemId       uint64   `json:"item_id"`
	ItemSku      string   `json:"item_sku"`
	VariationSku string   `json:"variation_sku"`
}

type GetReturnListResponse struct {
	BaseResponse

	Response GetReturnListResponseData `json:"response"`
}

func (s *ReturnServiceOp) GetReturnList(sid uint64, item GetReturnListRequest, tok string) (*GetReturnListResponse, error) {
	path := "/returns/get_return_list"

	resp := new(GetReturnListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, item)
	return resp, err
}
