package homecontroller

import (
	"html"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	println("access Home")
	w.Write([]byte(html.EscapeString("selamat datang ekoo")))
}
