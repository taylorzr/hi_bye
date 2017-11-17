package notify

import(
	"fmt"
	"strings"

	"github.com/taylorzr/hibye/hipchat"
	"github.com/taylorzr/hibye/root"
)

func Notify(fired []root.User, hired []root.User) (err error) {
	if len(fired) > 0 {
		message := buildMessage("Goodbye :(", fired)

		err = hipchat.SendMessage(message, hipchat.Red)

		if err != nil {
			return err
		}
	} else {
		// log.Println("No one fired since last run :)")
	}

	if len(hired) > 0 {
		message := buildMessage("Hello :)", hired)

		err = hipchat.SendMessage(message, hipchat.Yellow)

		if err != nil {
			return err
		}
	} else {
		// log.Println("No one hired since last run :/")
	}

	return nil
}

func buildMessage(header string, users []root.User) (message string) {
	messageLines := []string{ header }

	for _, user := range users {
		messageLines = append(messageLines, fmt.Sprintf("  - %s", user.Name))
	}

	return strings.Join(messageLines, "\n")
}
