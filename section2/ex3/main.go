package main

import (
	"log"
	"os"
	"text/template"
)

type dish struct {
	Name  string
	Price float64
}

type meal struct {
	Meal   string
	Dishes []dish
}

type menu []meal

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := menu{
		meal{
			Meal: "Breakfast",
			Dishes: []dish{
				dish{
					Name:  "Porrige",
					Price: 1.25,
				},
				dish{
					Name:  "Coffee",
					Price: 1.5,
				},
			},
		},
		meal{
			Meal: "Lunch",
			Dishes: []dish{
				dish{
					Name:  "Soup",
					Price: 2.0,
				},
				dish{
					Name:  "Steak",
					Price: 25.0,
				},
			},
		},
		meal{
			Meal: "Dinner",
			Dishes: []dish{
				dish{
					Name:  "Coffee",
					Price: 1.6,
				},
				dish{
					Name:  "Cupcake",
					Price: 1.2,
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", m)
	if err != nil {
		log.Fatalln(err)
	}
}
