package xfers

// Charge represent the xfers charge information
type Charge struct {
	ID                  string  `json:"id"`
	CheckoutURL         string  `json:"checkout_url"`
	NotifyURL           string  `json:"notify_url"`
	ReturnURL           string  `json:"return_url"`
	CancelURL           string  `json:"cancel_url"`
	Object              string  `json:"object"` // 'charge'
	Amount              float64 `json:"amount,string"`
	Currency            string  `json:"currency"`
	Customer            string  `json:"customer"`
	OrderID             string  `json:"order_id"`
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
