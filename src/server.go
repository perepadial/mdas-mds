package server

//Test pipeline
import (
	"fmt"
	"net/http"
)

func MainPageContent() string {
	return "Hello World ! \n"
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, MainPageContent())

}

func Init() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}
