package main

import (
	nav "go-nav/component"
	"os"
	"text/template"
)

func main() {
	templateCollection := template.Must(template.ParseGlob("templates/base.tmpl"))

	//bring in "external" nav component
	//and register it by its name
	templateCollection.AddParseTree(nav.GetComponentName(), nav.GetTemplates().Tree)

	println(templateCollection.DefinedTemplates())

	//base now renders component:nav
	templateCollection.ExecuteTemplate(os.Stdout, "base.tmpl", nav.GetDefaultData())

}
