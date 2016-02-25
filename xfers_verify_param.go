package xfers

// VerifyParam is used to VerifyCharge
type VerifyParam struct {
	OrderID     string  `json:"order_id"`
	TotalAmount float64 `json:"total_amount,string"`
	Currency    string  `json:"currency"`
	Status      string  `json:"status"`
}

// Init a helper function to convert NotifyParam into VerifyParam
func (x *VerifyParam) Init(param NotifyParam) {
	x.OrderID = param.OrderID
	x.TotalAmount = param.TotalAmount
	x.Currency = param.Currency
	x.Status = param.Status
}
