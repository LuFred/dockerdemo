package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/hi", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		w.Write([]byte("nihao"))
	})
	if err := http.ListenAndServe("127.0.0.1:5000", router); err != nil {
		log.Panicln(err)
	}
}
