package main

import (
	"log"
	"os"
	"text/template"
)

type region int

const (
	// Southern region
	Southern region = iota

	// Central region
	Central

	// Nothern region
	Nothern
)

func (r region) ToString() string {
	switch r {
	case Southern:
		return "Southern"
	case Central:
		return "Central"
	case Nothern:
		return "Nothern"
	default:
		return "Unknown"
	}
}

type hotel struct {
	Name, Address, City, Zip string
	Region                   region
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	califoniaHotels := []hotel{
		hotel{
			Name:    "Deluxe H",
			Address: "1, Boardwalk str.",
			City:    "Los Angeles",
			Zip:     "12345",
			Region:  Nothern,
		},
		hotel{
			Name:    "Awesome Stays Inc.",
			Address: "7, Heaven Blvd.",
			City:    "Santa Barbara",
			Zip:     "54321",
			Region:  Southern,
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", califoniaHotels)
	if err != nil {
		log.Fatalln(err)
	}
}
