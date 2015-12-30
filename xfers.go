package xfers

import (
	"bytes"
	"encoding/json"
	"errors"
	// "fmt"
	"io/ioutil"
	"net/http"
	// "net/url"
)

const XFERS_ENDPOINT_SANDBOX = "https://sandbox.xfers.io/api/v3"
const XFERS_ENDPOINT = "https://www.xfers.io/api/v3"

type XfersClient struct {
	Endpoint  string
	key       string
	isSandbox bool
	client    *http.Client
}

type XfersError struct {
	Error string
}

type XfersMsg struct {
	Msg string `json:"msg"`
}

type XfersAccount struct {
	AvailableBalance float64
	LedgerBalance    float64
	CreditCardRate   float64
	CreditCardFee    string
	BankTransferFee  float64
	FirstName        string
	LastName         string
	AddressLine1     string
	AddressLine2     string
	Nationality      string
	PostalCode       string
	IdentityNo       string
	Country          string
	Email            string
	IdBack           string
	IdDocument       string
	IdFront          string
	IdSelfie         string
	PhoneNo          string
}

type XfersChargeReqParam struct {
	Amount      string `json:"amount"`      // ! total XfersCharge.Items.price must sum up to XferCharge.Amount
	Currency    string `json:"currency"`    // !
	OrderId     string `json:"order_id"`    // !
	Description string `json:"description"` // !
	NotifyUrl   string `json:"notify_url"`
	ReturnUrl   string `json:"return_url"`
	CancelUrl   string `json:"cancel_url"`
	Refundable  string `json:"refundable"` // Default true
	// UserApiToken     string `json:"user_api_token"` // Optional
	// UserPhoneNo      string `json:"user_phone_no"`  // Default false
	// DebitOnly        string `json:"debit_only"`     // Default false
	Redirect string `json:"redirect"` // Default true
	// Items            string `json:"items"`          // JSON formatted array of items- description, name, price, quantity
	// Shipping         string `json:"shipping"`
	// Tax              string `json:"tax"`
	HrsToExpirations string `json:"hrs_to_expirations"` // Default to 48 hours from now
	// MetaData         string `json:"meta_data"`          // Key value pairs json
	ReceiptEmail string `json:"receipt_email"`
}

type XfersItem struct {
	Description string
	Name        string
	Price       float64
	Quantity    float64
	ItemId      string
}

type XfersCharge struct {
	Id                  string `json:"id"`
	CheckoutUrl         string `json:"checkout_url"`
	NotifyUrl           string `json:"notify_url"`
	ReturnUrl           string `json:"return_url"`
	CancelUrl           string `json:"cancel_url"`
	Object              string `json:"object"` // 'charge'
	Amount              string `json:"amount"`
	Currency            string `json:"currency"`
	Customer            string `json:"customer"`
	OrderId             string `json:"order_id"`
	Capture             bool   `json:"capture"`
	Refundable          bool   `json:"refundable"`
	Description         string `json:"description"`
	StatementDescriptor string `json:"statement_Descriptor"`
	ReceiptEmail        string `json:"receipt_email"`
	Shipping            string `json:"shipping"`
	Tax                 string `json:"tax"`
	TotalAmount         string `json:"total_amount"`
	Status              string `json:"status"`
}

type XfersVerifyParam struct {
	OrderId     string `json:"order_id"`
	TotalAmount string `json:"total_amount"`
	Currency    string `json:"currency"`
	Status      string `json:"status"`
}

func NewClient(key string, usesSandbox bool) (*XfersClient, error) {
	if key == "" {
		return nil, errors.New("Missing API Key")
	}
	xfersClient := XfersClient{}
	xfersClient.isSandbox = usesSandbox
	xfersClient.key = key
	xfersClient.client = new(http.Client)
	if xfersClient.isSandbox {
		xfersClient.Endpoint = XFERS_ENDPOINT_SANDBOX
	} else {
		xfersClient.Endpoint = XFERS_ENDPOINT
	}
	return &xfersClient, nil
}

func (xClient *XfersClient) PerformRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("X-XFERS-USER-API-KEY", xClient.key)
	req.Header.Add("Content-Type", "application/json")

	resp, err := xClient.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%v", string(body))
	if body == nil {
		return nil, errors.New("Missing body")
	}

	if string(body[0:1]) != "[" {
		errorResponse := XfersError{}
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			return nil, err
		}
		if errorResponse.Error != "" {
			return nil, errors.New(errorResponse.Error)
		}
	}

	return body, err
}

func (xClient *XfersClient) GetAccountInfo() (XfersAccount, error) {
	xfersAccount := XfersAccount{}
	req, err := http.NewRequest("GET", xClient.Endpoint+"/user", nil)
	if err != nil {
		return xfersAccount, err
	}
	resp, err := xClient.PerformRequest(req)
	if err != nil {
		return xfersAccount, err
	}
	err = json.Unmarshal(resp, &xfersAccount)
	return xfersAccount, err
}

func (xClient *XfersClient) CreateCharge(param XfersChargeReqParam) (XfersCharge, error) {
	xfersCharge := XfersCharge{}
	paramJson, err := json.Marshal(param)
	if err != nil {
		return xfersCharge, err
	}
	req, err := http.NewRequest("POST", xClient.Endpoint+"/charges", bytes.NewReader(paramJson))
	if err != nil {
		return xfersCharge, err
	}
	resp, err := xClient.PerformRequest(req)
	if err != nil {
		return xfersCharge, err
	}
	err = json.Unmarshal(resp, &xfersCharge)
	return xfersCharge, err
}

func (xClient *XfersClient) RetrieveCharge(id string) (XfersCharge, error) {
	xfersCharge := XfersCharge{}
	req, err := http.NewRequest("GET", xClient.Endpoint+"/charges/"+id, nil)
	if err != nil {
		return xfersCharge, err
	}
	resp, err := xClient.PerformRequest(req)
	if err != nil {
		return xfersCharge, err
	}
	err = json.Unmarshal(resp, &xfersCharge)
	return xfersCharge, err
}

func (xClient *XfersClient) ListAllCharges() ([]XfersCharge, error) {
	xfersCharges := []XfersCharge{}
	// supported params
	// customer - only returns specified customer ID
	// ending_before - cursor for position, use order_id
	// starting_after
	// limit
	req, err := http.NewRequest("GET", xClient.Endpoint+"/charges", nil)
	if err != nil {
		return xfersCharges, err
	}
	resp, err := xClient.PerformRequest(req)
	if err != nil {
		return xfersCharges, err
	}
	err = json.Unmarshal(resp, &xfersCharges)
	return xfersCharges, err
}

func (xClient *XfersClient) VerifyCharge(id string, param XfersVerifyParam) (XfersMsg, error) {
	xfersMsg := XfersMsg{}
	paramJson, err := json.Marshal(param)
	if err != nil {
		return xfersMsg, err
	}
	req, err := http.NewRequest("POST", xClient.Endpoint+"/charges/"+id+"/validate", bytes.NewReader(paramJson))
	if err != nil {
		return xfersMsg, err
	}
	resp, err := xClient.PerformRequest(req)
	if err != nil {
		return xfersMsg, err
	}
	err = json.Unmarshal(resp, &xfersMsg)
	return xfersMsg, err
}
