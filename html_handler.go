package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func doGreet(w http.ResponseWriter, r *http.Request) {
	// Parse template
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err%v", err)
		return
	}
	// Render template
	name := "Wenliang"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("Execute template failed, err%v", err)
		return
	}
}
