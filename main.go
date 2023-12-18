package main 

import (
  //"fmt"
  "net/http"
  "html/template"
  //"github.com/microcosm-cc/bluemonday"
)

func main() {

  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("app/static"))))
  http.HandleFunc("/", handleApp)
  http.HandleFunc("/render", renderHTML)
  // Start HTTP server
  http.ListenAndServe(":6969", nil)
}

func handleApp(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("app/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
    
    err = t.ExecuteTemplate(w, "app", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
}


func renderHTML(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
  w.Write([]byte(code))
}

