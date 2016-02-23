package xfers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	// EndpointSandbox Xfers API endpoint for sandbox / testing
	EndpointSandbox = "https://sandbox.xfers.io/api/v3"

	// Endpoint Xfers API endpoint for production
	Endpoint = "https://www.xfers.io/api/v3"

	// Fee Xfers Commission Fee
	Fee = 0.89 // as of 29 Jan 2016
)

// Client The main Xfers client interface
type Client struct {
	Endpoint  string
	key       string
	isSandbox bool
	client    *http.Client
}

// NewClient Initialize Xfers Client
func NewClient(key string, usesSandbox bool) (*Client, error) {
	if key == "" {
		return nil, errors.New("Missing API Key")
	}
	xClient := Client{}
	xClient.isSandbox = usesSandbox
	xClient.key = key
	xClient.client = new(http.Client)
	if xClient.isSandbox {
		xClient.Endpoint = EndpointSandbox
	} else {
		xClient.Endpoint = Endpoint
	}
	return &xClient, nil
}

// PerformRequest is a Helper function to perform api call to Xfers server
// and provide common exception checking
func (xClient *Client) PerformRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("X-XFERS-USER-API-KEY", xClient.key)
	req.Header.Add("Content-Type", "application/json")

	resp, err := xClient.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if body == nil {
		return nil, errors.New("Missing body")
	}

	if string(body[0:1]) != "[" {
		errorResponse := Error{}
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

// GetAccountInfo get the account info of the specified Xfers Client
func (xClient *Client) GetAccountInfo() (Account, error) {
	xfersAccount := Account{}
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

// CreateCharge create a new Xfers Charge
func (xClient *Client) CreateCharge(param ChargeReqParam) (Charge, error) {
	xfersCharge := Charge{}
	param.Redirect = false // Force to JSON Response!
	paramJSON, err := json.Marshal(param)
	if err != nil {
		return xfersCharge, err
	}
	req, err := http.NewRequest("POST", xClient.Endpoint+"/charges", bytes.NewReader(paramJSON))
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

// RetrieveCharge retrieve a particular charge
func (xClient *Client) RetrieveCharge(id string) (Charge, error) {
	xfersCharge := Charge{}
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

// ListAllCharges list last 10 created charges
func (xClient *Client) ListAllCharges() ([]Charge, error) {
	xfersCharges := []Charge{}
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

// VerifyCharge verify the notification call from the Xfers server
// Make sure the values in the request is valid
func (xClient *Client) VerifyCharge(id string, param VerifyParam) (Msg, error) {
	xfersMsg := Msg{}
	paramJSON, err := json.Marshal(param)
	if err != nil {
		return xfersMsg, err
	}
	req, err := http.NewRequest("POST", xClient.Endpoint+"/charges/"+id+"/validate", bytes.NewReader(paramJSON))
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
