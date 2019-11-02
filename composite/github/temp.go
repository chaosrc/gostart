package github

import (
	"fmt"
	html "html/template"
	"os"
	"text/template"
)

func Text(result *IssuesResult) {
	text := `
		count issues: {{.TotalCount}}
		{{range .Items}}-----
		number: {{.Number}}
		user: {{.User.Login}}
		title: {{.Title}}
		{{end}}
	`
	temp, err := template.New("result").Parse(text)
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("./out")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	temp.Execute(f, result)
}

func HTML(result *IssuesResult) {
	t := `
	<html>
		<head>
		</head>
		<body>
		<h3>Total{{.TotalCount}}</h3>
		<ul>
			{{range .Items}}
			<li>
			{{.Title}}  {{.Number}} <a href={{.URL}}>{{.User.Login}}</a>
			</li>
			{{end}}
		</ul>
		</body>
	</html>
	`
	h, err := html.New("issues").Parse(t)
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Create("./out.html")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	err2 := h.Execute(f, result)
	if err != nil {
		fmt.Println(err2)
	}
}
