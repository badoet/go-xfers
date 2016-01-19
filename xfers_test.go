package xfers_test

import (
	"../go-xfers"
	"crypto/rand"
	"encoding/binary"
	// "os"
	// "strings"
	"fmt"
	random "math/rand"
	"net/http"
	"testing"
)

const XFERS_ENDPOINT_SANDBOX = "https://sandbox.xfers.io/api/v3"
const XFERS_ENDPOINT = "https://www.xfers.io/api/v3"
const IS_SANDBOX = true
const TEST_KEY = "DqwoaCmbbLCLAYkZZ5DxgrwBxFdK_srHQn7XSTiBnWk"
const RECEIPT_EMAIL = "gunawan.stanley@gmail.com"

func TestNewClient(t *testing.T) {
	_, err := xfers.NewClient("", true)
	if err == nil {
		t.Errorf("Expected an error, due to missing API Key.")
	}

	sandboxClient, _ := xfers.NewClient(TEST_KEY, true)
	if sandboxClient.Endpoint != XFERS_ENDPOINT_SANDBOX {
		t.Errorf("Expected a sandbox endpoint, but get: %v", sandboxClient.Endpoint)
	}

	productionClient, _ := xfers.NewClient(TEST_KEY, false)
	if productionClient.Endpoint != XFERS_ENDPOINT {
		t.Errorf("Expected a production endpoint, but get: %v", productionClient.Endpoint)
	}
}

func TestPerformRequest(t *testing.T) {
	fmt.Println("Xfers: Perform Request")
	xClient, _ := xfers.NewClient(TEST_KEY, IS_SANDBOX)
	req, _ := http.NewRequest("GET", xClient.Endpoint+"/authorize/hello", nil)
	response, err := xClient.PerformRequest(req)
	if err != nil {
		t.Errorf("Did not expect any error, but get: %s", err.Error())
	}
	fmt.Printf("%s", response)
	fmt.Println("")
}

func TestGetAccount(t *testing.T) {
	fmt.Println("Xfers: Get Account")
	xClient, _ := xfers.NewClient(TEST_KEY, IS_SANDBOX)
	account, err := xClient.GetAccountInfo()
	if err != nil {
		t.Errorf("Did not expect any error, but get: %s", err.Error())
	}
	fmt.Printf("%+v", account)
	fmt.Println("")
}

func TestCreateCharge(t *testing.T) {
	fmt.Println("Xfers: Create Charge")
	xClient, _ := xfers.NewClient(TEST_KEY, IS_SANDBOX)

	chargeParam := xfers.XfersChargeReqParam{}
	chargeParam.Amount = "100.23"
	chargeParam.Shipping = "0.0"
	chargeParam.Currency = "SGD"
	chargeParam.OrderId = RandSeq(10)
	chargeParam.Description = "Test create charge"
	chargeParam.NotifyUrl = "https://www.ikoustyle.com/test/xfers/callback"
	chargeParam.ReturnUrl = "http://test.com/return"
	chargeParam.CancelUrl = "http://test.com/cancel"
	chargeParam.Redirect = "false"
	chargeParam.Refundable = "true"
	chargeParam.ReceiptEmail = RECEIPT_EMAIL
	chargeParam.HrsToExpirations = "48.0"

	chargeParam.Items = []xfers.XfersItem{}

	item1 := xfers.XfersItem{}
	item1.Description = "Item 1 Test"
	item1.Price = 20.23
	item1.Quantity = 1
	item1.Name = "Item 1"

	item2 := xfers.XfersItem{}
	item2.Description = "Item 2 Test"
	item2.Price = 80
	item2.Quantity = 1
	item2.Name = "Item 2"

	// NOTE: item1.Price + item2.Price must === chargeParam.Amount!

	chargeParam.Items = append(chargeParam.Items, item1, item2)

	xfersCharge, err := xClient.CreateCharge(chargeParam)
	if err != nil {
		t.Errorf("Did not expect any error, but get: %s", err.Error())
		return
	}
	fmt.Println("Xfers Charge:")
	fmt.Printf("%+v", xfersCharge)
	fmt.Println("")

	fmt.Println("Xfers: Retrieve Charge Details")
	xfersChargeCheck, err := xClient.RetrieveCharge(xfersCharge.Id)
	if err != nil {
		t.Errorf("Did not expect any error, but get: %s", err.Error())
	}
	fmt.Println("Xfers Charge Check:")
	fmt.Printf("%+v", xfersChargeCheck)
	fmt.Println("")
}

func TestListAllCharges(t *testing.T) {
	fmt.Println("Xfers: List All Charges")
	xClient, _ := xfers.NewClient(TEST_KEY, IS_SANDBOX)
	xfersCharges, err := xClient.ListAllCharges()
	if err != nil {
		t.Errorf("Did not expect any error, but get: %s", err.Error())
	}
	fmt.Printf("%+v", xfersCharges)
	fmt.Println("")
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ3456789")
var poolSize = len(letters)

func RandSeq(n int) string {
	var bb int64
	binary.Read(rand.Reader, binary.BigEndian, &bb)
	random.Seed(bb)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[random.Intn(poolSize)]
	}
	return string(b)
}
