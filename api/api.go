package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ApiServer struct {
	mux *httprouter.Router
}

func (a *ApiServer) Start(addr string) {
	a.mux = httprouter.New()
	a.mux.GET("/", a.Index)
	err := http.ListenAndServe(addr, a.mux)
	if err != nil {
		log.Fatalf("Error starting api server %s", err)
	}
}

func (a *ApiServer) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
