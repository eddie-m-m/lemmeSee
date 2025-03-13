package main

import (
	"html/template"
	"os"
)

type Hobby struct {
	Name template.HTML
}
type User struct {
	Name    template.HTML
	Age     int
	Hobbies []Hobby
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name:    "Dude",
		Age:     50,
		Hobbies: []Hobby{{Name: "Reading"}, {Name: "Coding"}},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
