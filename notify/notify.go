package notify

import(
	"fmt"
	// "io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/taylorzr/hibye/hipchat"
	"github.com/taylorzr/hibye/root"
	"github.com/taylorzr/hibye/storage"
)

const messageHourLimit = 24

func Notify(fired root.Users, hired root.Users) (err error) {
	if len(fired) > 0 {
		err := hipchat.Send(hipchat.Notification{
			Message: buildMessage("Goodbye :(", fired),
			Color: hipchat.Red,
			Notify: true,
		})

		if err != nil {
			return err
		}

		UpdateTimestamp()
	}

	if len(hired) > 0 {
		err := hipchat.Send(hipchat.Notification{
			Message: buildMessage("Hello :)", hired),
			Color: hipchat.Green,
			Notify: true,
		})

		if err != nil {
			return err
		}

		UpdateTimestamp()
	}

	if !notifiedRecently() {
		if len(fired) == 0 && len(hired) == 0 {
			err := hipchat.Send(hipchat.Notification{
				Message: "No one recently fired or hired :)",
				Color: hipchat.Green,
			})

			if err != nil {
				return err
			}

			UpdateTimestamp()
		}
	}

	return nil
}

func notifiedRecently() bool {
	t, err := storage.ReadMessageTimestamp()

	if err != nil {
		log.Fatal(err)
	}

	duration := time.Since(t)

	return duration.Hours() < messageHourLimit
}

func UpdateTimestamp() {
	err := storage.WriteMessageTimestamp(time.Now())

	if err != nil {
		log.Fatal(err)
	}
}

func buildMessage(header string, users root.Users) (message string) {
	messageLines := []string{ header }

	for _, user := range users {
		messageLines = append(messageLines, fmt.Sprintf("  - %s", user.Name))
	}

	return strings.Join(messageLines, "\n")
}
