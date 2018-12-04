package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type record struct {
	Date time.Time
	Open float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", table)
	http.ListenAndServe(":8887", nil)
}

func table(response http.ResponseWriter, request *http.Request) {
	records := parse("table.csv")

	err := tpl.ExecuteTemplate(response, "tpl.gohtml", records)
	if err != nil {
		log.Fatalln(err)
	}
}

func parse(path string) []record {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()

	records := make([]record, 0, len(rows)-1)

	for _, r := range rows[1:] {
		date, _ := time.Parse("2006-01-02", r[0])
		open, _ := strconv.ParseFloat(r[1], 64)

		records = append(records, record{
			Date: date,
			Open: open,
		})
	}

	return records
}
