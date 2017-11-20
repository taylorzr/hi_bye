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

const(
	Yellow = Color("yellow")
	Green  = Color("green")
	Red    = Color("red")
	Purple = Color("purple")
	Gray   = Color("gray")
	Random = Color("random")

	room = "hibye"
)

func SendMessage(message string, options Options) error {
	options = options.withDefaults()

	log.Println(message)

	client := new(http.Client)

	url := fmt.Sprintf("https://api.hipchat.com/v2/room/%s/notification?auth_token=%s", room, token)

	json, err := json.Marshal(map[string]interface{}{ 
		"message": message,
		"message_format": options.Format,
		"color": options.Color,
		"notify": options.DontNotify,
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

type Options struct {
	Color      Color
	DontNotify bool
	Format     string
}

func (options *Options) withDefaults() Options {
	if options.Color == "" {
		options.Color = Yellow
	}

	if options.Format == "" {
		options.Format = "text"
	}

	return *options
}
