package input

type UserInput struct {
	UserName  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Photo     string `json:"photo"`
	Telephone string `json:"telephone"`
}

type UserEditInput struct {
	Id        int    `json:"id" binding:"required"`
	UserName  string `json:"username" binding:"required"`
	Photo     string `json:"photo"`
	Telephone string `json:"telephone"`
}

type ChangePasswordInput struct {
	Id       int    `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
