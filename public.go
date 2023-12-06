package goshopee

type PublicService interface {
	GetShopsByPartner() (*GetShopsByPartnerResponse, error)
}

type PublicServiceOp struct {
	client *Client
}

type AuthedShop struct {
	Region     string `json:"region"`
	ShopId     uint64 `json:"shopid"`
	AuthTime   uint64 `json:"auth_time"`
	ExpireTime uint64 `json:"expire_time"`
}

type GetShopsByPartnerResponse struct {
	BaseResponse

	AuthedShops []AuthedShop `json:"authed_shop_list"`
}

func (s *PublicServiceOp) GetShopsByPartner() (*GetShopsByPartnerResponse, error) {
	path := "/public/get_shops_by_partner"

	resp := new(GetShopsByPartnerResponse)
	err := s.client.Get(path, resp, nil)
	return resp, err
}
