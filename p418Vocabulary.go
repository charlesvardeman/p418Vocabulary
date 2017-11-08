package main

import (
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

// MyServer struct for mux router
type MyServer struct {
  r *mux.Router
}

func main() {

  htmlRouter := mux.NewRouter().StrictSlash(true)
  htmlRouter.HandleFunc("/voc/diagram", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/diagram.html")
  })
  htmlRouter.HandleFunc("/voc/examples", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/examples.html")
  })
  htmlRouter.HandleFunc("/voc/examples/minimal", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/static/schema/examples/required.jsonld")
  })
  htmlRouter.HandleFunc("/voc/examples/full", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/static/schema/examples/full.jsonld")
  })
  htmlRouter.HandleFunc("/voc/schema.rdf", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/static/schema/schema.owl")
  })
  htmlRouter.HandleFunc("/voc/schema.xml", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/static/schema/schema.owl")
  })
  htmlRouter.HandleFunc("/voc/schema.owl", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/static/schema/schema.owl")
  })
  htmlRouter.HandleFunc("/voc/schema.json", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/static/schema/schema.jsonld")
  })
  htmlRouter.HandleFunc("/voc/schema.jsonld", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/static/schema/schema.jsonld")
  })
  htmlRouter.HandleFunc("/voc/schema.jsonld", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/static/schema/schema.jsonld")
  })
  htmlRouter.HandleFunc("/voc/schema.ttl", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/static/schema/schema.ttl")
  })
  htmlRouter.HandleFunc("/voc/schema.n3", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/static/schema/schema.n3")
  })
  htmlRouter.PathPrefix("/voc/static").Handler(http.StripPrefix("/voc/static", http.FileServer(http.Dir("./html/voc/static"))))
  htmlRouter.HandleFunc("/voc/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/index.html")
  })
  htmlRouter.HandleFunc("/voc/{resource}", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/voc/index.html")
  })
  http.Handle("/", &MyServer{htmlRouter})

  err := http.ListenAndServe(":9900", nil)
  // http 2.0 http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
  if err != nil {
    log.Fatal(err)
  }
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
  rw.Header().Set("Access-Control-Allow-Origin", "*")
  rw.Header().Set("Access-Control-Allow-Methods", "POST, GET")
  rw.Header().Set("Access-Control-Allow-Headers",
    "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

  // Let the Gorilla work
  s.r.ServeHTTP(rw, req)
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    fn(w, r)
  }
}
