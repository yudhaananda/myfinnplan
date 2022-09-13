package input

type BankAccountInput struct {
	AccountCode      string  `json:"accountcode" binding:"required"`
	AccountNameOwner string  `json:"accountnameowner" binding:"required"`
	BankCode         string  `json:"bankcode" binding:"required"`
	Amount           float64 `json:"amount" binding:"required"`
	Notes            string  `json:"notes" binding:"required"`
}

type BankAccountEditInput struct {
	Id               int     `json:"id" binding:"required"`
	AccountCode      string  `json:"accountcode" binding:"required"`
	AccountNameOwner string  `json:"accountnameowner" binding:"required"`
	BankCode         string  `json:"bankcode" binding:"required"`
	Amount           float64 `json:"amount" binding:"required"`
	Notes            string  `json:"notes" binding:"required"`
}
