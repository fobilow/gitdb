package main

import "text/template"

var packageTmpl = template.Must(template.New("package").Parse(`package gitdb
// Code generated by gitdb embed-ui on {{.Date}}; DO NOT EDIT.

func init() {
	//Embed Files
	{{range .Files}}
	getFs().embed("{{.Name}}", "{{.Content}}")
	{{end}}
}
`))
