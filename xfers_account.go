package xfers

// Account represent the xfers client account info
type Account struct {
	AvailableBalance float64 `json:"available_balance,string"`
	LedgerBalance    float64 `json:"ledger_balance,string"`
	CreditCardRate   float64 `json:"credit_card_rate,string"`
	CreditCardFee    string  `json:"credit_card_fee"`
	BankTransferFee  float64 `json:"bank_transfer_fee,string"`
	FirstName        string  `json:"first_name"`
	LastName         string  `json:"last_name"`
	AddressLine1     string  `json:"address_line1"`
	AddressLine2     string  `json:"address_line2"`
	Nationality      string  `json:"nationality"`
	PostalCode       string  `json:"postal_code"`
	IdentityNo       string  `json:"identity_no"`
	Country          string  `json:"country"`
	Email            string  `json:"email"`
	IDBack           string  `json:"id_back"`
	IDDocument       string  `json:"id_document"`
	IDFront          string  `json:"id_front"`
	IDSelfie         string  `json:"id_selfie"`
	PhoneNo          string  `json:"phone_now"`
}
