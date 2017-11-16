package main

import(
	"fmt"
	"os"
	"log"
	"strings"

	"github.com/taylorzr/hibye/root"
	"github.com/taylorzr/hibye/hipchat"
	"github.com/taylorzr/hibye/storage"
	"github.com/fatih/set"
)

func compare(oldUsers []root.User, newUsers []root.User) map[string][]root.User {
	oldSet := idSet(oldUsers)
	newSet := idSet(newUsers)

	firedIDs := set.IntSlice(set.Difference(oldSet, newSet))
	hiredIDs := set.IntSlice(set.Difference(newSet, oldSet))

	firedUsers := findUsers(firedIDs, oldUsers)
	hiredUsers := findUsers(hiredIDs, newUsers)

	return map[string][]root.User{
		"fired": firedUsers,
		"hired": hiredUsers,
	}
}

func findUsers(ids []int, users []root.User) []root.User {
	foundUsers := []root.User{}

	if len(ids) > 0 {
		usersByID := map[int]root.User{}

		for _, user := range users {
			usersByID[user.ID] = user
		}

		for _, id := range ids {
			foundUsers = append(foundUsers, usersByID[id])
		}
	}

	return foundUsers
}

func idSet(users []root.User) *set.Set {
	idSet := set.New()

	for _, user := range users {
		idSet.Add(user.ID)
	}

	return idSet
}

func notify(result map[string][]root.User) (err error) {
	if len(result["fired"]) > 0 {
		message := buildMessage("Goodbye :(", result["fired"])

		err = hipchat.SendMessage(message, hipchat.Red)

		if err != nil {
			return err
		}
	} else {
		log.Println("No one fired since last run :)")
	}

	if len(result["hired"]) > 0 {
		message := buildMessage("Hello :)", result["hired"])

		err = hipchat.SendMessage(message, hipchat.Yellow)

		if err != nil {
			return err
		}
	} else {
		log.Println("No one hired since last run :/")
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

func notmain() {
	log.Printf("Hitting up hipchat for all the users...")

	users, err := hipchat.GetAllUsers()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d users!", len(users))

	err = storage.Write("old_users.csv", users)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Wrote user data to users.csv")
}

func oldmain() {
	oldUsers, err := storage.Read("old_users.csv")

	if err != nil {
		log.Fatal(err)
	}

	newUsers, err := storage.Read("new_users.csv")

	if err != nil {
		log.Fatal(err)
	}

	result := compare(oldUsers, newUsers)

	err = notify(result)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if _, err := os.Stat("old_users.csv"); !os.IsNotExist(err) {
		oldUsers, err := storage.Read("old_users.csv")
		if err != nil {
			log.Fatal(err)
		}

		newUsers, err := hipchat.GetAllUsers()
		if err != nil {
			log.Fatal(err)
		}

		defer storage.Write("old_users.csv", newUsers)

		comparison := compare(oldUsers, newUsers)

		err = notify(comparison)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		newUsers, err := hipchat.GetAllUsers()
		if err != nil {
			log.Fatal(err)
		}

		err = storage.Write("old_users.csv", newUsers)

		if err != nil {
			log.Fatal(err)
		}

		err = hipchat.SendMessage("HiBye initialized, will report on next run", hipchat.Yellow)

		if err != nil {
			log.Fatal(err)
		}
	}
}
