package input

type UserInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserEditInput struct {
	Id       int    `json:"id" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
