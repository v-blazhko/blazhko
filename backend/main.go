package main

import (
	"github.com/v-blazhko/blazhko/api"
	"github.com/v-blazhko/blazhko/controller"
	"net/http"
	"os"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3000"
	}

	c := controller.Controller{}
	handler := api.Handler(&c)

	_ = http.ListenAndServe(":"+PORT, handler)
}
