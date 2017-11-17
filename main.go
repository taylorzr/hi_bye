package main

import(
	"os"
	"log"

	"github.com/taylorzr/hibye/compare"
	"github.com/taylorzr/hibye/hipchat"
	"github.com/taylorzr/hibye/notify"
	"github.com/taylorzr/hibye/storage"
)


func main() {
	old_users_path := "old_users.csv"
	// old_users_path := "test_old_users.csv"

	if _, err := os.Stat(old_users_path); !os.IsNotExist(err) {
		oldUsers, err := storage.Read(old_users_path)
		if err != nil {
			log.Fatal(err)
		}

		newUsers, err := hipchat.GetAllUsers()
		// newUsers, err := storage.Read("test_new_users.csv")
		if err != nil {
			log.Fatal(err)
		}

		defer storage.Write(old_users_path, newUsers)

		hiredUsers := compare.FindFired(oldUsers, newUsers)
		firedUsers := compare.FindHired(oldUsers, newUsers)

		err = notify.Notify(hiredUsers, firedUsers)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		newUsers, err := hipchat.GetAllUsers()
		if err != nil {
			log.Fatal(err)
		}

		err = storage.Write(old_users_path, newUsers)

		if err != nil {
			log.Fatal(err)
		}

		err = hipchat.SendMessage("HiBye initialized, will report on next run", hipchat.Yellow)

		if err != nil {
			log.Fatal(err)
		}
	}
}
