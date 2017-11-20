package hipchat

import(
	"os"
	"log"
)

var token, hipchatTokenExists = os.LookupEnv("HIPCHAT_TOKEN")

func init() {
	if !hipchatTokenExists {
		log.Fatal("HIPCHAT_TOKEN environment variable not set")
	}
}
