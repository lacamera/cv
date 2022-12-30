package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	var output string
	var data string
	flag.StringVar(&output, "o", "resume.html", "")
	flag.StringVar(&data, "c", "resume.yaml", "")
	flag.Parse()
	tmpl := os.Args[len(os.Args)-1]

	values := map[string]interface{}{}
	byt, err := os.ReadFile(data)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(byt, &values)
	if err != nil {
		log.Fatal(err)
	}

	byt, err = os.ReadFile(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	t := template.Must(template.New("").Parse(string(byt)))
	err = t.Execute(f, values)
	if err != nil {
		log.Fatal(err)
	}
}
