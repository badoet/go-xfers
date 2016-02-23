package xfers

// NotifyParam is the structure of the data sent by Xfers server
// to notify and update order status to your server
type NotifyParam struct {
	TxnId       string  `form:"txn_id" json:"txn_id" binding:"required"`
	OrderId     string  `form:"order_id" json:"order_id" binding:"required"`
	TotalAmount float64 `form:"total_amount" json:"total_amount" binding:"required"`
	Currency    string  `form:"currency" json:"currency" binding:"required"`
	Status      string  `form:"status" json:"status" binding:"required"` // “cancelled” or “paid” or “expired”
}
