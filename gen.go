package main

import (
	"html/template"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	data := map[string]interface{}{}
	buf, err := os.ReadFile("data.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		log.Fatal(err)
	}
	byt, err := os.ReadFile("resume.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("resume.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	t := template.Must(template.New("").Parse(string(byt)))
	err = t.Execute(f, data)
	if err != nil {
		log.Fatal(err)
	}
}
