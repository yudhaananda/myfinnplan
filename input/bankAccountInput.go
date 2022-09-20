package input

type BankAccountInput struct {
	UserAccountId  int     `json:"useraccountid" binding:"required"`
	AccountIdOwner int     `json:"accountidowner" binding:"required"`
	BankCode       string  `json:"bankcode" binding:"required"`
	Amount         float64 `json:"amount" binding:"required"`
	Notes          string  `json:"notes" binding:"required"`
	IsDebit        bool    `json:"isdebit"`
}

type BankAccountEditInput struct {
	Id             int     `json:"id" binding:"required"`
	UserAccountId  int     `json:"useraccountid" binding:"required"`
	AccountIdOwner int     `json:"accountidowner" binding:"required"`
	BankCode       string  `json:"bankcode" binding:"required"`
	Amount         float64 `json:"amount" binding:"required"`
	Notes          string  `json:"notes" binding:"required"`
	IsDebit        bool    `json:"isdebit"`
}
