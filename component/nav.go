package nav

import (
	"embed"
	"text/template"
)

//go:embed templates/*.tmpl
var Templates embed.FS

type NavItem struct {
	Title    string
	Link     string
	Selected bool
}

type NavList struct {
	Title string
	Items []NavItem
}

func GetTemplates() *template.Template {
	tmpl := template.Must(template.ParseFS(Templates, "templates/nav.tmpl"))
	return tmpl
}

func GetComponentName() string {
	id := "component:nav"
	return id
}

func GetDefaultData() NavList {
	nav := NavList{
		Title: "Page menu",
		Items: []NavItem{
			{Title: "Home", Link: "/", Selected: true},
			{Title: "Assets", Link: "/assets", Selected: false},
		},
	}
	return nav
}
