package main

import (
	"encoding/csv"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type closing struct {
	Date time.Time
	Open float32
}
type closings []closing

var tmpl *template.Template

var fm = template.FuncMap{
	"fdateDMY": formatDate,
}

func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.gohtml"))
}

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8080", nil)
}

func serve(res http.ResponseWriter, req *http.Request) {
	cl := parseCSV("table.csv")
	tmpl.ExecuteTemplate(res, "tmpl.gohtml", cl)
}

func formatDate(t time.Time) string {
	return t.Format("02.01.2006")
}

func parseCSV(table string) closings {

	f, err := os.Open(table)
	if err != nil {
		log.Fatalln("Could not open file!")
	}
	defer f.Close()

	reader := csv.NewReader(f)

	columnNames, err := reader.Read()
	lineCount := 1
	cl := closings{}
	log.Println("column names: ", columnNames)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("Error: ", err)
			break
		}
		date, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			log.Println("Error:", err)
			break
		}
		price, err := strconv.ParseFloat(record[4], 32)
		if err != nil {
			log.Println("Error:", err)
			break
		}

		closing := closing{
			date,
			float32(price),
		}
		cl = append(cl, closing)
		// fmt.Println("Record", lineCount, "has", len(record), "fields:", record)
		lineCount++
	}
	return cl
	// fmt.Println("Total Number of lines: ", lineCount)
	// fmt.Println(cl)
}
