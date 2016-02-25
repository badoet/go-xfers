package xfers

// Item is used to build XfersChargeParam
type Item struct {
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Quantity    float64 `json:"quantity"`
	ItemID      string  `json:"item_id,omitempty"` // Optional
}
