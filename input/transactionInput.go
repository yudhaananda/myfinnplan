package input

type TransactionInput struct {
	BankAccountId int     `json:"bankaccountid" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
	Notes         string  `json:"notes" binding:"required"`
}

type TransactionEditInput struct {
	Id            int     `json:"id" binding:"required"`
	BankAccountId int     `json:"bankaccountid" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
	Notes         string  `json:"notes" binding:"required"`
}
