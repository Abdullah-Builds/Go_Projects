package main

import (
	"bytes"
	"html/template"
	"sync"
)

type Receipent struct {
	Name  string
	Email string
}

func main() {

	recipient := make(chan Receipent)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		LoadFile("./demo_records.csv", recipient)
		defer wg.Done()
		close(recipient)
	}()

	go func() {
		EmailWorker(1, recipient)
		defer wg.Done()
	}()

	wg.Wait()
}


func ExecuteTemplate(r Receipent)(string,error){
	t,err := template.ParseFiles("email.tmpl")

	if err != nil{
		return "",err
	}

	var tpl bytes.Buffer

	err = t.Execute(&tpl,r)
	if err != nil{
		return "",err
	}
    return tpl.String(),nil

}