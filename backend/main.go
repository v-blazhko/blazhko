package main

import (
	"github.com/v-blazhko/blazhko/backend/api"
	"github.com/v-blazhko/blazhko/backend/controller"
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
