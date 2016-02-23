package xfers

// Account represent the xfers client account info
type Account struct {
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
