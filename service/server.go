package service

import (
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/unrolled/render"
)

func NewServer() *negroni.Negroni {
	formatter := render.New()
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	mx.HandleFunc("/", home).Methods("GET")
	mx.HandleFunc("/unknown", unknown).Methods("GET")
	mx.HandleFunc("/table", login).Methods("POST")
	mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
}
