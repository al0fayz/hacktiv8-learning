package rules

type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}
