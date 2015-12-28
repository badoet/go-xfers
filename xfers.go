package xfers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

type XfersCharge struct {
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

type XfersChargeResponse struct {
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

	// fmt.Println(xClient.Endpoint + apiUrl)
	// fmt.Println("X-XFERS-USER-API-KEY", xClient.key)

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

	errorResponse := XfersError{}
	err = json.Unmarshal(body, &errorResponse)
	if err != nil {
		return nil, err
	}
	if errorResponse.Error != "" {
		return nil, errors.New(errorResponse.Error)
	}

	return body, err
}

func (xClient *XfersClient) GetAccountInfo() (XfersAccount, error) {
	account := XfersAccount{}
	req, err := http.NewRequest("GET", xClient.Endpoint+"/user", nil)
	if err != nil {
		return account, err
	}
	resp, err := xClient.PerformRequest(req)
	if err != nil {
		return account, err
	}
	err = json.Unmarshal(resp, &account)
	return account, err
}

func (xClient *XfersClient) CreateCharge(charge XfersCharge) (XfersChargeResponse, error) {
	response := XfersChargeResponse{}
	// values := url.Values{}
	// values.Add("amount", fmt.Sprintf("%.2f", charge.Amount))
	// values.Add("currency", charge.Currency)
	// values.Add("order_id", charge.OrderId)
	// values.Add("description", charge.Description)
	// values.Add("notify_url", charge.NotifyUrl)
	// values.Add("return_url", charge.ReturnUrl)
	// values.Add("cancel_url", charge.CancelUrl)
	// values.Add("redirect", charge.Redirect)
	// values.Add("receipt_email", charge.ReceiptEmail)
	// req, err := http.NewRequest("POST", xClient.Endpoint+"/charges", bytes.NewBufferString(values.Encode()))
	chargeJson, _ := json.Marshal(charge)
	fmt.Printf("%s", string(chargeJson))
	req, err := http.NewRequest("POST", xClient.Endpoint+"/charges", bytes.NewReader(chargeJson))
	if err != nil {
		fmt.Println("Err create request")
		return response, err
	}
	resp, err := xClient.PerformRequest(req)
	fmt.Println("Create Charge Response:")
	fmt.Printf("%s\n", resp)
	if err != nil {
		fmt.Println("Err perform request")
		return response, err
	}
	err = json.Unmarshal(resp, &response)
	return response, err
}
