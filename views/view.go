package views

import (
	"fmt"
	"html/template"
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
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{Template: t, Layout: layout}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	if err := v.Template.ExecuteTemplate(w, v.Layout, data); err != nil {
		return err
	}
	return nil
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