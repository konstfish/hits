package models

type BadgeParams struct {
	URL     string `form:"url" binding:"required"`
	TitleBg string `form:"title_bg"`
	CountBg string `form:"count_bg"`
	Title   string `form:"title"`
}
