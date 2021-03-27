package views

import (
	"bytes"
	"fmt"
	"github.com/thiduzz/lenslocked.com/context"
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

func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	var vd Data
	switch d := data.(type) {
	case Data:
		vd = d
	default:
		vd = Data{
			Alert: nil,
			Yield: data,
		}
	}
	vd.User = context.User(r.Context())
	var bf bytes.Buffer
	if err := v.Template.ExecuteTemplate(&bf, v.Layout, vd); err != nil {
		http.Error(w, "Something went wrong, contact the admins...", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &bf)
}

func (v *View) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	v.Render(writer, request,nil)
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