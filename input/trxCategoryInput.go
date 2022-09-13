package input

type TrxCategoryInput struct {
	CategoryCode string `json:"categorycode" binding:"required"`
	CategoryName string `json:"categoryname" binding:"required"`
}

type TrxCategoryEditInput struct {
	Id           int    `json:"id" binding:"required"`
	CategoryCode string `json:"categorycode" binding:"required"`
	CategoryName string `json:"categoryname" binding:"required"`
}
