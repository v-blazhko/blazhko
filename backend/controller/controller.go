package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Controller struct {
}

func (c *Controller) PostApiContact(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Println(name, value)
		}
	}

	fmt.Print(body)
	_, _ = fmt.Fprintf(w, "Hello World from path: %s\n, message: %s", r.URL.Path, string(body))
}
