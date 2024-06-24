package main

import (
	"fmt"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {

	fileName:= "loading.html"

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("unable to read %s", fileName)
	}

	fmt.Printf("Request received\n")
	w.Write(data)
}

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}