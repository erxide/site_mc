package handle

import (
	"html/template"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("./templates/test.html"))
	page.Execute(w, r)
}
