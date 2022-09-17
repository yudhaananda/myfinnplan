package input

type UserAccountInput struct {
	AccountName string `json:"accountname" binding:"required"`
}

type UserAccountEditInput struct {
	Id          int    `json:"id" binding:"required"`
	AccountName string `json:"accountname" binding:"required"`
}
