package controller

import (
	"encoding/json"
	"github.com/v-blazhko/blazhko/backend/api"
	"github.com/v-blazhko/blazhko/backend/ds"
	"github.com/v-blazhko/blazhko/backend/mail"
	"io/ioutil"
	"log"
	"net/http"
)

type Controller struct {
	storage *ds.DS
}

func (c *Controller) PostApiContact(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("An error occurred while submitting your response"))
		return
	}

	contact := api.ClientContact{}

	err = json.Unmarshal(body, &contact)
	if err != nil {
		log.Printf("Error unmarshalling: %s, body: %s", err, body)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("An error occurred while submitting your response"))
		return
	}

	err = c.storage.CreateNewClientContact(contact)
	if err != nil {
		log.Printf("Error creating contact: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = mail.SendToMyself(contact.Message, contact.Name, contact.Email)
	if err != nil {
		log.Printf("Error sending email: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewController(storage *ds.DS) Controller {
	return Controller{
		storage: storage,
	}
}
