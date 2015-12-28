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

// const TEST_KEY = "DqwoaCmbbLCLAYkZZ5DxgrwBxFdK_srHQn7XSTiBnWk"

const TEST_KEY = "Jaf5Q4K7Ns7jSbEtaFyxNPtaDNZJ4EfNnDZx7XzUyyc"

func TestNewClient(t *testing.T) {
	_, err := xfers.NewClient("", true)
	if err == nil {
		t.Errorf("Expected an error, due to missing API Key.")
	}

	sandboxClient, _ := xfers.NewClient(TEST_KEY, true)
	if sandboxClient.Endpoint != XFERS_ENDPOINT_SANDBOX {
		t.Errorf("Expected a sandbox endpoint, but get: %v.", sandboxClient.Endpoint)
	}

	productionClient, _ := xfers.NewClient(TEST_KEY, false)
	if productionClient.Endpoint != XFERS_ENDPOINT {
		t.Errorf("Expected a production endpoint, but get: %v.", productionClient.Endpoint)
	}
}

func TestPerformRequest(t *testing.T) {
	fmt.Println("Xfers: Perform Request")
	xClient, _ := xfers.NewClient(TEST_KEY, false)
	req, _ := http.NewRequest("GET", xClient.Endpoint+"/authorize/hello", nil)
	response, err := xClient.PerformRequest(req)
	if err != nil {
		t.Errorf("Did not expect any error, but get: %s.", err.Error())
	}
	fmt.Printf("%s", response)
	fmt.Println("")
}

func TestGetAccount(t *testing.T) {
	fmt.Println("Xfers: Get Account")
	xClient, _ := xfers.NewClient(TEST_KEY, false)
	account, err := xClient.GetAccountInfo()
	if err != nil {
		t.Errorf("Did not expect any error, but get: %s.", err.Error())
	}
	fmt.Printf("%v", account)
	fmt.Println("")
}

func TestCreateCharge(t *testing.T) {
	fmt.Println("Xfers: Create Charge")
	xClient, _ := xfers.NewClient(TEST_KEY, false)
	charge := xfers.XfersCharge{}
	charge.Amount = "100.232"
	charge.Currency = "SGD"
	charge.OrderId = RandSeq(10)
	charge.Description = "Test create charge"
	charge.NotifyUrl = "http://test.com/notify"
	charge.ReturnUrl = "http://test.com/return"
	charge.CancelUrl = "http://test.com/cancel"
	charge.Redirect = "false"
	charge.Refundable = "true"
	charge.CancelUrl = "admin@email.com"
	charge.HrsToExpirations = "48.0"
	chargeResponse, err := xClient.CreateCharge(charge)
	if err != nil {
		t.Errorf("Did not expect any error, but get: %s.", err.Error())
	}
	fmt.Println("Charge resp")
	fmt.Printf("%v", chargeResponse)
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
