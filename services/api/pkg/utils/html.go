package utils

import (
	"bytes"
	"html/template"
	"time"
)

func templateHTML(x string) template.HTML { return template.HTML(x) }
func timeToStr(i int64) string            { return time.Unix(i, 0).Format("2006-01-02 15:04:05") }
func add(a, b int) int                    { return a + b }

// HTML 生成html
func HTML(html string, data interface{}) string {
	t := template.New("html").Delims("{{", "}}")
	t = t.Funcs(template.FuncMap{
		"unescaped": templateHTML,
		"timeToStr": timeToStr,
		"add":       add,
	})
	t, _ = t.Parse(html)
	bf := bytes.NewBuffer(nil)
	_ = t.Execute(bf, data)
	return bf.String()
}
