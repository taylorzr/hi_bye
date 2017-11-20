package notify

import(
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/taylorzr/hibye/hipchat"
	"github.com/taylorzr/hibye/root"
	"github.com/taylorzr/hibye/storage"
)

const messageHourLimit = 1
const last_message_path = "hibye_last_message"

func Notify(fired []root.User, hired []root.User) (err error) {
	if !storage.Exists(last_message_path) {
		updateTimestamp()
		// log.Println("Initialized last run time")
	}

	if len(fired) > 0 {
		message := buildMessage("Goodbye :(", fired)

		err = hipchat.SendMessage(message, hipchat.Options{ Color: hipchat.Red })

		if err != nil {
			return err
		}

		updateTimestamp()
	}

	if len(hired) > 0 {
		message := buildMessage("Hello :)", hired)

		err = hipchat.SendMessage(message, hipchat.Options{ Color: hipchat.Green })

		if err != nil {
			return err
		}

		updateTimestamp()
	}

	if !notifiedRecently() {
		if len(fired) == 0 && len(hired) == 0 {
			err = hipchat.SendMessage("No one recently fired or hired :)", hipchat.Options{ Color: hipchat.Green })

			if err != nil {
				return err
			}

			updateTimestamp()
		}
	}

	return nil
}

func notifiedRecently() bool {
	timeData, err := ioutil.ReadFile(last_message_path)

	if err != nil {
		log.Fatal(err)
	}

	t := time.Time{}

	err = t.UnmarshalText(timeData)

	if err != nil {
		log.Fatal(err)
	}

	duration := time.Since(t)

	return duration.Hours() < messageHourLimit
}

func updateTimestamp() {
	timeText, _ := time.Now().MarshalText()

	err := ioutil.WriteFile(last_message_path, []byte(timeText), 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func buildMessage(header string, users []root.User) (message string) {
	messageLines := []string{ header }

	for _, user := range users {
		messageLines = append(messageLines, fmt.Sprintf("  - %s", user.Name))
	}

	return strings.Join(messageLines, "\n")
}
