package handler

import (
	"github.com/gin-gonic/gin"
	ui "github.com/konstfish/ui/core"
	"github.com/konstfish/ui/themes/kf"
)

// gin handler index
func HandleIndex(c *gin.Context) {
	page := ui.NewPage().
		SetTitle("konstfish/hits").
		AddScript("https://unpkg.com/htmx.org@2.0.4").
		AddScript("https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/prism.min.js").
		AddScript("https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/components/prism-go.min.js").
		AddScript("https://cdnjs.cloudflare.com/ajax/libs/jquery/3.7.1/jquery.min.js").
		AddStyleSheet("https://cdn.jsdelivr.net/gh/konstfish/ui@main/static/main.css").
		AddStyleSheet("https://cdn.jsdelivr.net/gh/konstfish/ui@main/static/gallery/etc.css").
		AddStyleSheet("https://cdn.jsdelivr.net/gh/konstfish/ui@main/static/prism.css").
		AddLinkWithType("icon", "static/logo.svg", "image/svg+xml")

	page.Body.AddChild(kf.HeaderBar(kf.TitleLogo("konstfish/hits", "static/logo.svg"), []kf.KeyValue{{Key: "Source", Value: "https://github.com/konstfish/hits"}}))

	main := kf.AppBody().AddChild(kf.Panel(kf.TitleLogo("", "/api/count/incr/badge.svg?url=https://hits.konst.fish&count_bg=%236580A8&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits")).AddClasses("flex-center"))
	main.AddChild(kf.Panel(kf.Code("text", "GET https://hits.konst.fish/api/count/incr/badge.svg?url=https://hits.konst.fish&count_bg=#6580A8&title_bg=#555555&icon=&icon_color=#E7E7E7&title=hits")))

	page.Body.AddChild(main)

	out, err := page.Render()
	if err != nil {
		c.String(500, "oops")
		c.Header("Refresh", "2")
		return
	}

	c.Writer.Write([]byte(out))
}
