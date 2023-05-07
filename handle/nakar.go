package handle

import (
	"html/template"
	"net/http"
)

func Nakar(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("./templates/nakar.html"))
	page.Execute(w, r)
}
