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
type restaurant struct {
	Name string
	Menu meals
}
type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	restaurants := restaurants{
		{
			"The Small Buck",
			meals{
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
			},
		},
		{
			"The Big Bucket",
			meals{
				{
					"Biggest 1",
					"Breakfast",
					33448.8,
					items{
						{"lotsa egg"},
						{"lotsa onion"},
					},
				},
				{
					"Big night",
					"Dinner",
					1934.3,
					items{
						{"lotsa steak"},
						{"buckets of sauce"},
					},
				},
			},
		},
	}
	tpl.Execute(os.Stdout, restaurants)
}
