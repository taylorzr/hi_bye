package storage

import(
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/taylorzr/hi_bye/root"
)

func Write(path string, users []root.User) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)

	defer writer.Flush()

	for _, user := range users {
		data := []string{ fmt.Sprintf("%d", user.ID), user.Name }

		err := writer.Write(data)

		if err != nil {
			return err
		}
	}

	return nil
}

func Read(path string) ([]root.User, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)

	userData, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	users := []root.User{}

	for _, userDatum := range userData {
		id, err := strconv.Atoi(userDatum[0])
		if err != nil { return nil, err }

		users = append(users, root.User{ID: id, Name: userDatum[1]})
	}

	return users, nil
}
