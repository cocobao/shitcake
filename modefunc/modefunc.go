package modefunc

import "html/template"

var (
	FuncsMap template.FuncMap = template.FuncMap{
		"addcount": AddCount,
	}
)

func AddCount(s int, c int) int {
	return (s + c)
}
