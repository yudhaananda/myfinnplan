package input

type TrxCategoryInput struct {
	CategoryName string `json:"categoryname" binding:"required"`
}

type TrxCategoryEditInput struct {
	Id           int    `json:"id" binding:"required"`
	CategoryName string `json:"categoryname" binding:"required"`
}
