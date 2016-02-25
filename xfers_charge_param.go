package xfers

// ChargeReqParam is used to create Xfers Charge
type ChargeReqParam struct {
	Amount           float64 `json:"amount,string"` // ! total XfersCharge.Items.price must sum up to XferCharge.Amount
	Currency         string  `json:"currency"`      // !
	OrderID          string  `json:"order_id"`      // !
	Description      string  `json:"description"`   // !
	NotifyURL        string  `json:"notify_url,omitempty"`
	ReturnURL        string  `json:"return_url,omitempty"`
	CancelURL        string  `json:"cancel_url,omitempty"`
	Refundable       bool    `json:"refundable,omitempty,string"` // Default true
	Redirect         bool    `json:"redirect,string"`             // Default true
	Items            []Item  `json:"items,omitempty"`             // JSON formatted array of items- description, name, price, quantity
	Shipping         float64 `json:"shipping,omitempty,string"`
	Tax              float64 `json:"tax,omitempty,string"`
	HrsToExpirations float64 `json:"hrs_to_expirations,omitempty,string"` // Default to 48 hours from now
	ReceiptEmail     string  `json:"receipt_email,omitempty"`
	UserAPIToken     string  `json:"user_api_token,omitempty"`       // Optional
	UserPhoneNo      bool    `json:"user_phone_no,omitempty,string"` // Default false
	DebitOnly        bool    `json:"debit_only,omitempty,string"`    // Default false
	MetaData         string  `json:"meta_data,omitempty"`            // Key value pairs json
}
