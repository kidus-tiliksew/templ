// Code generated by templ@(devel) DO NOT EDIT.

package testcssexpression

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "strings"

func className() templ.CSSClass {
	var templCSSBuilder strings.Builder
	templCSSBuilder.WriteString(`background-color:#ffffff;`)
	templCSSBuilder.WriteString(string(templ.SanitizeCSS(`color`, red)))
	templCSSID := templ.CSSID(`className`, templCSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templCSSID,
		Class: templ.SafeCSS(`.` + templCSSID + `{` + templCSSBuilder.String() + `}`),
	}
}
