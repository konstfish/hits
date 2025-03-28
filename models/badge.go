package models

type BadgeParams struct {
	URL     string `form:"url" binding:"required"`
	CountBg string `form:"count_bg"`
	Title   string `form:"title"`
}

func (bp *BadgeParams) SetDefaults() error {
	if bp.CountBg == "" {
		bp.CountBg = "#007ec6"
	}
	if bp.Title == "" {
		bp.Title = "hits"
	}

	return nil
}
