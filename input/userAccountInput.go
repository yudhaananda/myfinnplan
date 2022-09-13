package input

type UserAccountInput struct {
	AccountCode string `json:"accountcode" binding:"required"`
	AccountName string `json:"accountname" binding:"required"`
}

type UserAccountEditInput struct {
	Id          int    `json:"id" binding:"required"`
	AccountCode string `json:"accountcode" binding:"required"`
	AccountName string `json:"accountname" binding:"required"`
}
