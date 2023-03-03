package route

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World") 	
}

func Bye(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bye World") 	
}



