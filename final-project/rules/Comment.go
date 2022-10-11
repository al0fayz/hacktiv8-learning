package rules

type CommentCreate struct {
	Message string `form:"message" json:"message" binding:"required"`
}
