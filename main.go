package main

import(
	"fmt"
	"log"

	"github.com/taylorzr/hi_bye/root"
	"github.com/taylorzr/hi_bye/hipchat"
	"github.com/taylorzr/hi_bye/storage"
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

func notMain() {
	oldUsers, _ := storage.Read("old_users.csv")
	newUsers, _ := storage.Read("new_users.csv")

	result := compare(oldUsers, newUsers)

	if len(result["fired"]) > 0 {
		fmt.Println("Goodbye :(")

		for _, user := range result["fired"] {
			fmt.Printf("  - %s\n", user.Name)
		}
	}

	if len(result["hired"]) > 0 {
		fmt.Println("Hello :)")

		for _, user := range result["hired"] {
			fmt.Printf("  - %s\n", user.Name)
		}
	}
}

func main() {
	err := hipchat.SendMessage("Howdy")

	if err != nil {
		log.Fatal(err)
	}
}
