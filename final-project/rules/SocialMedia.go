package rules

type SocialMediaCreate struct {
	Name           string `form:"name" json:"name" binding:"required"`
	SocialMediaUrl string `form:"social_media_url" json:"social_media_url" binding:"required"`
}
