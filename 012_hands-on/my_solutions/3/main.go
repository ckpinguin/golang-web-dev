package main

import "html/template"
import "os"

var tpl *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  region
}
type region struct {
	Name string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.go.html"))
}
func main() {
	regions := []region{
		region{"Southern"},
		region{"Central"},
		region{"Northern"},
	}
	california := []hotel{
		hotel{
			`Frisco star`,
			`Upper st. 4`,
			`San Francisco`,
			4444,
			regions[1]},
		hotel{
			`Southern Star`,
			`Lower downtown`,
			`Southolo`,
			5555,
			regions[2]},
	}
	tpl.Execute(os.Stdout, california)
}
