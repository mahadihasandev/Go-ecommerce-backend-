package main

import (
	"fmt"
	"net/http"
)

func handleAbout(write http.ResponseWriter, read *http.Request) {
	fmt.Fprintln(write, "this is About")
}