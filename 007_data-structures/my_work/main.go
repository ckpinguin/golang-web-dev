package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("006_variable/my_work/templ.tmpl"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "templ.tmpl", `J E S U S`)
	if err != nil {
		log.Fatalln(err)
	}
}
