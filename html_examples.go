package main

import (
	"log"
	"os"
	"text/template"
)

type Part struct {
	Name string
	Count int
}

type Subscriber struct {
	Name string
	Rate float64
	Active bool
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	templateText := "Name {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	subscriber := Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	executeTemplate(templateText, subscriber)
	subscriber = Subscriber{Name: "Joy Carr", Rate: 5.99, Active: false}
	executeTemplate(templateText, subscriber)
	
	// templateText := "Name: {{.Name}}\nCount: {{.Count}}\n"
	// executeTemplate(templateText, Part{Name: "Fuses", Count: 5})
	// executeTemplate(templateText, Part{Name: "Cables", Count: 2})

	// templateText := "Before loop: {{.}}\n{{range .}}In llop: {{.}}\n{{end}}After loop: {{.}}\n"
	// executeTemplate(templateText, []string{"do", "re", "mi"})
	// templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	// executeTemplate(templateText, []float64{1.25, 0.99, 27})

	// templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	// executeTemplate(templateText, []float64{})
	// executeTemplate(templateText, nil)

	// executeTemplate("Dot is: {{if.}} Do is true!{{end}} finish\n", true)
	// executeTemplate("Dot is: {{if.}} Do is true!{{end}} finish\n", true)

	// templateText := "Template start\nAction: {{.}}\nTemaplate end\n"
	// tmpl, err := template.New("test").Parse(templateText)
	// check(err)
	// err = tmpl.Execute(os.Stdout, "ABC")
	// check(err)
	// err = tmpl.Execute(os.Stdout, 42)
	// check(err)
	// err = tmpl.Execute(os.Stdout, true)
	// check(err)

	// text := "Here's my template!\n"
	// tmpl, err := template.New("test").Parse(text)
	// check(err)
	// err = tmpl.Execute(os.Stdout, nil)
	// check(err)
}