package main

import(
	"log"
	"os"
	"fmt"
	"time"

	"github.com/taylorzr/hibye/compare"
	"github.com/taylorzr/hibye/hipchat"
	"github.com/taylorzr/hibye/notify"
	"github.com/taylorzr/hibye/storage"
)

const ignoreUsersFetchErrorLimitHours = 1

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	if !storage.Exists() {
		initialize()
	} else {
		check()
	}
}

func initialize() {
    storage.Create()
	notify.UpdateTimestamp()

	newUsers, err := hipchat.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}

	err = storage.WriteUsers(newUsers)
	if err != nil {
		log.Fatal(err)
	}

	err = hipchat.Send(hipchat.Notification{ Message: "HiBye initialized, will report on next run" })

	if err != nil {
		log.Fatal(err)
	}
}

func check() {
	oldUsers, err := storage.ReadUsers()
	if err != nil {
		log.Fatal(err)
	}

	newUsers, err := hipchat.GetAllUsers()
	if err != nil {
		lastUsersTime, err := storage.ReadUsersTimestamp()

		if err != nil {
			log.Fatal(err)
		}

		duration := time.Since(lastUsersTime)

		// TODO: This could also be handled by maybe checking if there is a
		// network connection and just not doing anything
		if duration.Hours() < ignoreUsersFetchErrorLimitHours {
			os.Exit(0)
		} else {
			log.Fatal(fmt.Sprintf("Error fetching hipchat users for the last %d hours", ignoreUsersFetchErrorLimitHours))
		}
	}

	defer storage.WriteUsers(newUsers)

	firedUsers := compare.FindFired(oldUsers, newUsers)
	hiredUsers := compare.FindHired(oldUsers, newUsers)

	err = notify.Notify(firedUsers, hiredUsers)

	if err != nil {
		log.Fatal(err)
	}
}

