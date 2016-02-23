package xfers

// Charge represent the xfers charge information
type Charge struct {
	Id                  string  `json:"id"`
	CheckoutUrl         string  `json:"checkout_url"`
	NotifyUrl           string  `json:"notify_url"`
	ReturnUrl           string  `json:"return_url"`
	CancelUrl           string  `json:"cancel_url"`
	Object              string  `json:"object"` // 'charge'
	Amount              float64 `json:"amount,string"`
	Currency            string  `json:"currency"`
	Customer            string  `json:"customer"`
	OrderId             string  `json:"order_id"`
	Capture             bool    `json:"capture"`
	Refundable          bool    `json:"refundable"`
	Description         string  `json:"description"`
	StatementDescriptor string  `json:"statement_Descriptor"`
	ReceiptEmail        string  `json:"receipt_email"`
	Shipping            float64 `json:"shipping,string"`
	Tax                 float64 `json:"tax,string"`
	TotalAmount         float64 `json:"total_amount,string"`
	Status              string  `json:"status"`
}
