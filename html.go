package tonic

import (
	"html/template"

	"github.com/gin-gonic/gin"

	"github.com/sknv/tonic/render"
	xtemplate "github.com/sknv/tonic/x/html/template"
)

func LoadHtmlWalk(engine *gin.Engine, templateRoot, templateExt string) {
	if gin.IsDebugging() {
		engine.HTMLRender = render.HtmlDebug{
			FuncMap:      engine.FuncMap,
			TemplateExt:  templateExt,
			TemplateRoot: templateRoot,
		}
		return
	}

	tpl := template.Must(
		xtemplate.ParseWalk(
			template.New(""), templateRoot, templateExt,
		),
	)
	engine.SetHTMLTemplate(tpl)
}
