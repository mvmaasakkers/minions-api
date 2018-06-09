package hackathon_api

type BankAccount struct {
	Id           string `json:"id" bson:"_id"`
	UserID       string `json:"user_id" bson:"user_id"`
	Label        string `json:"label"`
	BankID       string `json:"bank_id"`
	Transactions int    `json:"transactions"`
	AccountRouting struct {
		Scheme  string `json:"scheme"`
		Address string `json:"address"`
	} `json:"account_routing"`
}

type BankAccounts struct {
	BankAccounts []BankAccount `json:"accounts"`
}
