package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
)

type View struct {
	Template *template.Template
	Layout   string
}

var (
	LayoutDir string = "views/layouts"
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles(layout)...)
	files = append(files, componentFiles(layout)...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{Template: t, Layout: layout}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	switch data.(type) {
	case Data:
	default:
		data = Data{
			Alert: nil,
			Yield: data,
		}
	}
	var bf bytes.Buffer
	if err := v.Template.ExecuteTemplate(&bf, v.Layout, data); err != nil {
		http.Error(w, "Something went wrong, contact the admins...", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &bf)
}

func (v *View) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	v.Render(writer, nil)
}


func layoutFiles(layout string) []string {
	files, err := filepath.Glob(fmt.Sprintf("%s/%s/*%s",LayoutDir,layout, TemplateExt))
	if err != nil{
		panic(err)
	}
	return files
}

func componentFiles(layout string) []string {
	files, err := filepath.Glob(fmt.Sprintf("%s/%s/*%s",LayoutDir,layout, TemplateExt))
	if err != nil{
		panic(err)
	}
	return files
}

func addTemplatePath(files []string)  {
	for i, s := range files {
		files[i] = TemplateDir + s
	}
}

func addTemplateExt(files []string)  {
	for i, s := range files {
		files[i] = s + TemplateExt
	}
}