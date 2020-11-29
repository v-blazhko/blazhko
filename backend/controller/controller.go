package controller

import (
	"encoding/json"
	"fmt"
	"github.com/v-blazhko/blazhko/backend/api"
	"github.com/v-blazhko/blazhko/backend/ds"
	"github.com/v-blazhko/blazhko/backend/mail"
	"io/ioutil"
	"net/http"
)

type Controller struct {
	storage *ds.DS
}

func (c *Controller) PostApiContact(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("An error occurred while submitting your response"))
		return
	}

	contact := api.ClientContact{}

	err = json.Unmarshal(body, &contact)
	if err != nil {
		fmt.Printf("Error unmarshalling: %s, body: %s", err, body)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("An error occurred while submitting your response"))
		return
	}

	err = c.storage.CreateNewClientContact(contact)
	if err != nil {
		fmt.Printf("Error creating contact: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("An error occurred while submitting your response"))
		return
	}

	err = mail.SendToMyself(contact.Message, contact.Name, contact.Email)
	if err != nil {
		fmt.Printf("Error sending email: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("An error occurred while submitting your response"))
		return
	}

	_, _ = w.Write([]byte("Thanks for submitting your message! I will reach out to you as fast as possible :)"))
}

func NewController(storage *ds.DS) Controller {
	return Controller{
		storage: storage,
	}
}

func getErrorMessage(lang string) string {
	if lang == "ru" {
		return fmt.Sprint("Произошла ошибка! Попробуйте снова :/")
	} else {
		return fmt.Sprint("An error occurred while submitting your response")
	}
}

func getSuccessMessage(lang string) string {
	if lang == "ru" {
		return fmt.Sprint("Спасибо за ваше сообщение! Скоро я свяжусь с вами :)")
	} else {
		return fmt.Sprint("Thanks for submitting your message! I will reach out to you as fast as possible :)")
	}
}
