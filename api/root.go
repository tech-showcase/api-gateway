package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Activate(port int) {
	address := fmt.Sprintf(":%d", port)

	r := mux.NewRouter()
	RegisterMovieHTTPAPI(r)

	err := http.ListenAndServe(address, r)
	if err != nil {
		panic(err)
	}
}
