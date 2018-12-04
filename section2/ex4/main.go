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

type restaurant struct {
	Name string
	Menu menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := []restaurant{
		restaurant{
			Name: "Diner",
			Menu: menu{
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
			},
		},
		restaurant{
			Name: "Burger Truck",
			Menu: menu{
				meal{
					Meal: "Lunch",
					Dishes: []dish{
						dish{
							Name:  "Burger",
							Price: 4.25,
						},
					},
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", r)
	if err != nil {
		log.Fatalln(err)
	}
}
