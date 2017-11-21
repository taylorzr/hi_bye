package hipchat

import(
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Color string
type Format string

const(
	Yellow = Color("yellow")
	Green  = Color("green")
	Red    = Color("red")
	Purple = Color("purple")
	Gray   = Color("gray")
	Random = Color("random")

	HTML = Format("html")
	Text = Format("text")

	room = "hibye"
)

type Notification struct {
	Message string
	Color Color
	Format Format
	DontNotify bool
}

func (notification *Notification) withDefaults() Notification {
	if notification.Color == "" {
		notification.Color = Yellow
	}

	if notification.Format == "" {
		notification.Format = Text
	}

	return *notification
}

func Send(notification Notification) error {
	notification = notification.withDefaults()

	log.Println(notification.Message)

	client := new(http.Client)

	url := fmt.Sprintf("https://api.hipchat.com/v2/room/%s/notification?auth_token=%s", room, token)

	json, err := json.Marshal(map[string]interface{}{ 
		"message": notification.Message,
		"message_format": notification.Format,
		"color": notification.Color,
		"notify": notification.DontNotify,
	})

	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(json))

	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != 204 {
		message := fmt.Sprintf("Sending message: Expected 204 response, got %d", response.StatusCode)
		return errors.New(message)
	}

	return nil
}
