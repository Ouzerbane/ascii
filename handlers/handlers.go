package hand

import (
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"

	"youmed/ascii-art"
	check "youmed/utils"
)

type donner struct {
	Input   string
	Banner  string
	Status  int
	Message string
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Error(http.StatusMethodNotAllowed, w)
		return
	}
	if r.URL.Path != "/" {
		Error(http.StatusNotFound, w)
		return
	}
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		Error(http.StatusInternalServerError, w)
		return
	}
	tmpl.Execute(w, nil)
}

func Handlerascii(w http.ResponseWriter, r *http.Request) {
	v := ""

	if r.Method != http.MethodPost {
		Error(http.StatusMethodNotAllowed, w)
		return
	}
	tmpl, err := template.ParseFiles("templates/ascii-art.html")
	if err != nil {
		Error(http.StatusInternalServerError, w)
		return
	}
	data := donner{
		Input:  r.FormValue("input"),
		Banner: r.FormValue("banner"),
	}
	if (data.Input == "") || (data.Banner != "standard" && data.Banner != "shadow" && data.Banner != "thinkertoy") || !check.CheckIn(data.Input) {
		Error(http.StatusBadRequest, w)
		return
	}
	v = ascii.Ascii(data.Banner, data.Input)
	tmpl.Execute(w, v)
}

func Error(status int, w http.ResponseWriter) {
	tmp, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "enternal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	newdonner := donner{
		Status:  status,
		Message: http.StatusText(status),
	}
	tmp.Execute(w, newdonner)
}

func Static(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/static/" || (!strings.HasSuffix(r.URL.Path, "style.css") && !strings.HasSuffix(r.URL.Path, "a.png")) {
		Error(http.StatusNotFound, w)
		return
	}
	data, _ := os.ReadFile(strings.TrimPrefix(r.URL.Path, "/"))
	w.Header().Set("content-type", "text/css")
	w.Write(data)
}

func Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Error(http.StatusMethodNotAllowed, w)
		return
	}

	asciiArt := r.FormValue("Download")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Lenght", strconv.Itoa(len(asciiArt)))
	w.Write([]byte(asciiArt))
}
