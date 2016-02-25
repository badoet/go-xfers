package xfers_test

import (
	"../go-xfers"
	"crypto/rand"
	"encoding/binary"
	random "math/rand"
	"net/http"
	"os"
	"testing"
)

const isSandbox = true // for testing always use the sandbox version

func fetchEnvVars(t *testing.T) (key, notifyUrl string) {
	key = os.Getenv("XFERS_TEST_KEY")
	if len(key) <= 0 {
		t.Fatalf("Test cannot run because cannot get environment variable XFERS_TEST_KEY")
	}
	notifyUrl = os.Getenv("XFERS_NOTIFY_URL")
	if len(notifyUrl) <= 0 {
		t.Fatalf("Test cannot run because cannot get environment variable XFERS_NOTIFY_URL")
	}
	return
}

func TestNewClient(t *testing.T) {
	key, _ := fetchEnvVars(t)
	_, err := xfers.NewClient("", true)
	if err == nil {
		t.Errorf("Expected an error, due to missing API Key.")
	}

	sandboxClient, _ := xfers.NewClient(key, true)
	if sandboxClient.Endpoint != xfers.EndpointSandbox {
		t.Errorf("Expected a sandbox endpoint, but get: %v", sandboxClient.Endpoint)
	}

	productionClient, _ := xfers.NewClient(key, false)
	if productionClient.Endpoint != xfers.Endpoint {
		t.Errorf("Expected a production endpoint, but get: %v", productionClient.Endpoint)
	}
}

func TestPerformRequest(t *testing.T) {
	key, _ := fetchEnvVars(t)
	xClient, _ := xfers.NewClient(key, isSandbox)
	req, _ := http.NewRequest("GET", xClient.Endpoint+"/authorize/hello", nil)
	_, err := xClient.PerformRequest(req)
	if err != nil {
		t.Errorf("Did not expect any error, but get: %s", err.Error())
	}
}

func TestGetAccount(t *testing.T) {
	key, _ := fetchEnvVars(t)
	xClient, _ := xfers.NewClient(key, isSandbox)
	_, err := xClient.GetAccountInfo()
	if err != nil {
		t.Errorf("Did not expect any error, but get: %s", err.Error())
	}
}

func TestCreateCharge(t *testing.T) {
	key, notifyUrl := fetchEnvVars(t)
	xClient, _ := xfers.NewClient(key, isSandbox)

	chargeParam := xfers.ChargeReqParam{}
	chargeParam.Amount = 100.23
	chargeParam.Shipping = 0.0
	chargeParam.Currency = "SGD"
	chargeParam.OrderID = RandSeq(10)
	chargeParam.Description = "Test create charge"
	chargeParam.NotifyURL = notifyUrl
	chargeParam.ReturnURL = "http://test.com/return"
	chargeParam.CancelURL = "http://test.com/cancel"
	chargeParam.Refundable = true
	chargeParam.ReceiptEmail = "test@email.com"
	chargeParam.HrsToExpirations = 48.0

	chargeParam.Items = []xfers.Item{}

	item1 := xfers.Item{}
	item1.Description = "Item 1 Test"
	item1.Price = 20.23
	item1.Quantity = 1
	item1.Name = "Item 1"

	item2 := xfers.Item{}
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

	_, err = xClient.RetrieveCharge(xfersCharge.ID)
	if err != nil {
		t.Errorf("Did not expect any error, but get: %s", err.Error())
	}
}

func TestListAllCharges(t *testing.T) {
	key, _ := fetchEnvVars(t)
	xClient, _ := xfers.NewClient(key, isSandbox)
	_, err := xClient.ListAllCharges()
	if err != nil {
		t.Errorf("Did not expect any error, but get: %s", err.Error())
	}
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
