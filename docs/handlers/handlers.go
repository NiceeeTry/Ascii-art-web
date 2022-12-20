package handlers

import (
	ascii "ascii-art-web/docs/asciiTransform"
	"net/http"
	"text/template"
)

type Data struct {
	UserInput   string
	ErrorNumber int
	ErrorText   string
}

var tmpl = template.Must(template.ParseGlob("docs/htmlTemplates/*.html"))

func MainPageHandler(writer http.ResponseWriter, request *http.Request) {
	var data Data

	if request.URL.Path != "/" {
		data.ErrorNumber = 404
		data.ErrorText = "Page not found"
		errorHandler(writer, request, &data)
		return
	}
	if request.Method != http.MethodGet {
		data.ErrorText = "Method not allowed"
		data.ErrorNumber = 405
		errorHandler(writer, request, &data)
		return
	}
	err := tmpl.ExecuteTemplate(writer, "index.html", data)
	if err != nil {
		data.ErrorNumber = 500
		data.ErrorText = "Internal Server Error"
		errorHandler(writer, request, &data)
		return
	}
}

func AsciiPage(writer http.ResponseWriter, request *http.Request) {
	var data Data
	if request.URL.Path != "/ascii" {
		data.ErrorNumber = 404
		data.ErrorText = "Page not found"
		errorHandler(writer, request, &data)
		return
	}
	if request.Method != http.MethodPost {
		data.ErrorText = "Method not allowed"
		data.ErrorNumber = 405
		errorHandler(writer, request, &data)
		return
	}
	text := request.FormValue("input")
	font := request.FormValue("font")
	// fmt.Fprintf(writer, text)
	if !ascii.CheckAlpha(text) {
		data.ErrorNumber = 400
		data.ErrorText = "Bad requestuest"
		errorHandler(writer, request, &data)
		return
	}
	output, err := ascii.AsciiArt(text, font)
	if err {
		data.ErrorNumber = 500
		data.ErrorText = "Internal Server Error"
		errorHandler(writer, request, &data)
		return
	}
	data.UserInput = output
	tmpl.ExecuteTemplate(writer, "index.html", data)
}

func errorHandler(w http.ResponseWriter, r *http.Request, data *Data) {
	w.WriteHeader(data.ErrorNumber)
	tmpl.ExecuteTemplate(w, "error.html", data)
}
