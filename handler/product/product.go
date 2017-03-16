package product

import "html/template"

// Product func map
// i.e. baseFuncMap["test"] = func() string{
//  return "message"
// }
func productsFuncMap(baseFuncMap template.FuncMap) template.FuncMap {
	return baseFuncMap
}
