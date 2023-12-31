package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_GetDiscountDetail(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/discount/get_discount", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_discount_resp.json")))

	var req GetDiscountRequest
	loadMockData("get_discount_req.json", &req)

	res, err := client.Discount.GetDiscount(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Discount.GetDiscountDetail error: %s", err)
	}

	t.Logf("Discount.GetDiscount: %#v", res)

	var expectedID uint64 = 665123666665499
	if res.Response.DiscountID != expectedID {
		t.Errorf("DiscountID returned %+v, expected %+v", res.Response.DiscountID, expectedID)
	}
}

func Test_GetDiscountList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/discount/get_discount_list", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_discount_list_resp.json")))

	var req GetDiscountListRequest
	loadMockData("get_discount_list_req.json", &req)

	res, err := client.Discount.GetDiscountList(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Discount.GetDiscountList error: %s", err)
	}

	t.Logf("Discount.GetDiscountList: %#v", res)

	var expectedID uint64 = 1000021581
	if res.Response.DiscountList[0].DiscountID != expectedID {
		t.Errorf("DiscountID returned %+v, expected %+v", res.Response.DiscountList[0].DiscountID, expectedID)
	}
}

func Test_AddDiscount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/add_discount", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("add_discount_resp.json")))

	var req AddDiscountRequest
	loadMockData("add_discount_req.json", &req)

	res, err := client.Discount.AddDiscount(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Discount.AddDiscount error: %s", err)
	}

	t.Logf("Discount.AddDiscount: %#v", res)

	var expectedID uint64 = 122131231
	if res.Response.DiscountID != expectedID {
		t.Errorf("DiscountID returned %+v, expected %+v", res.Response.DiscountID, expectedID)
	}
}

func Test_AddDiscountItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/add_discount_item", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("add_discount_item_resp.json")))

	var req AddDiscountItemRequest
	loadMockData("add_discount_item_req.json", &req)

	res, err := client.Discount.AddDiscountItem(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Discount.AddDiscountItem error: %s", err)
	}

	t.Logf("Discount.AddDiscountItem: %#v", res)

	var expected string = "discount.error_time"
	if res.Response.ErrorList[0].FailError != expected {
		t.Errorf("FailError returned %+v, expected %+v", res.Response.ErrorList[0].FailError, expected)
	}
}

func Test_DeleteDiscountItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/delete_discount_item", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("delete_discount_item_resp.json")))

	res, err := client.Discount.DeleteDiscountItem(shopID, 1000029882, 1776783, 1467683, accessToken)
	if err != nil {
		t.Errorf("Discount.DeleteDiscountItem error: %s", err)
	}

	t.Logf("Discount.DeleteDiscountItem: %#v", res)

	var expected string = "time error"
	if res.Response.ErrorList[0].FailMessage != expected {
		t.Errorf("FailMessage returned %+v, expected %+v", res.Response.ErrorList[0].FailMessage, expected)
	}
}

func Test_UpdateDiscountItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/update_discount_item", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("update_discount_item_resp.json")))

	var req UpdateDiscountItemRequest
	loadMockData("update_discount_item_req.json", &req)

	res, err := client.Discount.UpdateDiscountItem(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Discount.UpdateDiscountItem error: %s", err)
	}

	t.Logf("Discount.UpdateDiscountItem: %#v", res)

	var expected string = "time error"
	if res.Response.ErrorList[0].FailMessage != expected {
		t.Errorf("FailMessage returned %+v, expected %+v", res.Response.ErrorList[0].FailMessage, expected)
	}
}
