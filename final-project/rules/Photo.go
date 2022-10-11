package rules

type PhotoCreate struct {
	Title    string `form:"title" json:"title" binding:"required"`
	PhotoUrl string `form:"photo_url" json:"photo_url" binding:"required"`
}
