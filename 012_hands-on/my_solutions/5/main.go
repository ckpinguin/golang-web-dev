package main

import "html/template"
import "os"

type item struct {
	Name string
}
type items []item

type meal struct {
	Name, Type string
	Price      float32
	Items      items
}
type meals []meal

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	menu := meals{
		{
			"Small 1",
			"Breakfast",
			48.8,
			items{
				{"egg"},
				{"onion"},
			},
		},
		{
			"Big two",
			"Dinner",
			19.3,
			items{
				{"steak"},
				{"sauce"},
			},
		},
	}
	tpl.Execute(os.Stdout, menu)

}
