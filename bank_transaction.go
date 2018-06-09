package hackathon_api

import "time"

type BankTransactions struct {
	BankTransactions []BankTransaction `json:"transactions"`
}

type BankTransaction struct {
	Id           string                  `json:"id" bson:"_id"`
	UserID       string                  `json:"user_id" bson:"user_id"`
	ThisAccount  TransactionThisAccount  `json:"this_account" bson:"this_account"`
	Details      TransactionDetails      `json:"details"`
	Metadata     TransactionMetadata     `json:"metadata"`
	OtherAccount TransactionOtherAccount `json:"other_account" bson:"other_account"`
}

type TransactionThisAccount struct {
	Id string `json:"id"`
	BankRouting struct {
		Scheme  string `json:"scheme"`
		Address string `json:"address"`
	} `json:"bank_routing"`
	AccountRouting struct {
		Scheme  string `json:"scheme"`
		Address string `json:"address"`
	} `json:"account_routing"`
	Holders []struct {
		Name    string `json:"name"`
		IsAlias bool   `json:"is_alias"`
	} `json:"holders"`
}

type TransactionDetails struct {
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Posted      time.Time `json:"posted"`
	Completed   time.Time `json:"completed"`
	NewBalance struct {
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"new_balance"`
	Value struct {
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"value"`
}

type TransactionMetadata struct {
	Narrative struct {
		Narrative string `json:"narrative"`
	} `json:"narrative"`
	Tags []struct {
		Id    string    `json:"id"`
		Date  time.Time `json:"date"`
		Value string    `json:"value"`
	} `json:"tags"`
	Images []struct {
		Id    string          `json:"id"`
		Date  time.Time       `json:"date"`
		URL   string          `json:"URL"`
		Label string          `json:"label"`
		User  TransactionUser `json:"user"`
	} `json:"images"`
	Comments []struct {
		Id    string          `json:"id"`
		Value string          `json:"value"`
		Date  time.Time       `json:"date"`
		User  TransactionUser `json:"user"`
	} `json:"comments"`
}

type TransactionUser struct {
	Id          string `json:"id"`
	Provider    string `json:"provider"`
	DisplayName string `json:"display_name"`
}

type TransactionOtherAccount struct {
	Id string `json:"id"`
	Holder struct {
		Name    string `json:"name"`
		IsAlias bool   `json:"is_alias"`
	} `json:"holder"`
	BankRouting struct {
		Scheme  string `json:"scheme"`
		Address string `json:"address"`
	} `json:"bank_routing"`
	AccountRouting struct {
		Scheme  string `json:"scheme"`
		Address string `json:"address"`
	} `json:"account_routing"`
}
