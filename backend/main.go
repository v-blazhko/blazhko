package main

import (
	"github.com/v-blazhko/blazhko/backend/api"
	"github.com/v-blazhko/blazhko/backend/controller"
	"github.com/v-blazhko/blazhko/backend/ds"
	"net/http"
	"os"
)

func main() {
	var ServicePort string
	if ServicePort = os.Getenv("SERVICE_PORT"); ServicePort == "" {
		ServicePort = "3000"
	}

	db, _ := ds.NewDB()
	storage := ds.NewStorage(db)

	c := controller.NewController(storage)
	handler := api.Handler(&c)

	_ = http.ListenAndServe(":"+ServicePort, handler)
}
