package internal

import (
	"embed"
	"io"
	"log"
	"text/template"
)

// embed the template files
//go:embed templates/resource.tmpl
//go:embed templates/provider.tmpl
//go:embed templates/service.tmpl
var tmplFS embed.FS

type Resource struct {
	HasServices          bool
	PackageName          string
	Messages             []*Message
	ProviderSchemaFields []MessageField
}

type Message struct {
	Name    string
	GoName  string
	Fields  []MessageField
	IDField *MessageField
	Create  *Method
	Read    *Method
	Update  *Method
	Delete  *Method
	Exists  *Method

	ignore bool
}

type MessageField struct {
	ID           bool
	Name         string
	GoName       string
	Type         string
	GoType       string
	Optional     bool
	Required     bool
	Description  string
	Computed     bool
	ForceNew     bool
	Elem         *Schema
	ResourceElem string
	Deprecated   bool
	Sensitive    bool
	SubSchema    *Message
}

type Service struct {
	ServiceName string
	Methods     []Method
}

type Method struct {
	ServiceName   string
	Function      string
	Operation     string
	Resource      string
	ParameterType string
	ReturnType    string
}

type Provider struct {
	PackageName  string
	ProviderName string
	Services     []Service
	Resources    []Resource
}

type Schema struct {
	Type string
}

func WriteTemplate(tmpl string, wr io.Writer, data interface{}) {
	t := template.Must(template.ParseFS(tmplFS, tmpl)) // parse embeded template
	err := t.Execute(wr, data)                         // execute template and write to file

	if err != nil {
		log.Print(err)
	}
}
