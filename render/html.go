package render

import (
	"html/template"

	"github.com/gin-gonic/gin/render"

	xtemplate "github.com/sknv/tonic/x/html/template"
)

type HtmlDebug struct {
	FuncMap      template.FuncMap
	TemplateExt  string
	TemplateRoot string
}

// Instance implements gin's HTMLRender interface.
func (r HtmlDebug) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r.loadTemplate(),
		Name:     name,
		Data:     data,
	}
}

func (r HtmlDebug) loadTemplate() *template.Template {
	return template.Must(
		xtemplate.ParseWalk(
			template.New("").Funcs(r.FuncMap),
			r.TemplateRoot,
			r.TemplateExt,
		),
	)
}
