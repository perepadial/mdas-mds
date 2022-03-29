package server

//Test pipeline
import (
	"fmt"
	"net/http"
)

func MainPageContent() string {
	return "Hello Go! \n"
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", MainPageContent())

}

func Init() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		fmt.Printf("error occurred during server initialization ---> %s", err.Error())
	}
}
