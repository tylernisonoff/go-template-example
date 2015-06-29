package main

import (
	"text/template"
	"os"
)

type Inventory struct {
	FirstHost string
	FirstPort string
	SecondHost string
	SecondPort string
}
func main() {
	tmpl, err := template.ParseFiles("template.tmp")
	sweaters := Inventory{"wool", "port", "hi", "port"}
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil { panic(err) }
}
