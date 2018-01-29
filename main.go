package main

import(
	"log"

	"github.com/taylorzr/hibye/compare"
	"github.com/taylorzr/hibye/hipchat"
	"github.com/taylorzr/hibye/notify"
	"github.com/taylorzr/hibye/storage"
)

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
	// newUsers, err := storage.Read("test_new_users.csv")
	if err != nil {
		// TODO:
		// Sometimes get an error written to stderr like:
		// 2018/01/24 10:04:05 Get https://api.hipchat.com/v2/user?max-results=1000&auth_token=ZqDKnJaxyiK7wEScsmMW2ad6yHD81P13Tl10Kwjr: dial tcp: lookup api.hipchat.com: no such host
		// This is because the laptop has no network connection, it would be
		// nice to only log when a successful run of this program hasn't been
		// made in the last few hours or something
		log.Fatal(err)
	}

	defer storage.WriteUsers(newUsers)

	firedUsers := compare.FindFired(oldUsers, newUsers)
	hiredUsers := compare.FindHired(oldUsers, newUsers)

	err = notify.Notify(firedUsers, hiredUsers)

	if err != nil {
		log.Fatal(err)
	}
}

