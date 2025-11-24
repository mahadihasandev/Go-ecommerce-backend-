package handlers

import (
	"fmt"
	"net/http"
)

func HandleAbout(write http.ResponseWriter, read *http.Request) {
	fmt.Fprintln(write, "this is About")
}