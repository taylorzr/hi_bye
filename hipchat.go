package main

import(
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"net/http"
)

var token = os.Getenv("HIPCHAT_TOKEN")

type UserResponse struct {
	Items []User `json:"items"`
}

type User struct {
	ID          int `json:"id"`
	Name        string `json:"name"`
	MentionName string `json:"mention_name"`
}

func GetAllUsers() ([]User, error) {
	httpBody, err := makeRequest()

	if err != nil {
		return nil, err
	}

	users, err := decodeResponse(httpBody)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func makeRequest() ([]byte, error) {
	client := new(http.Client)

	url := fmt.Sprintf("https://api.hipchat.com/v2/user?max-results=1000&auth_token=%s", token)

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	httpResponse, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func decodeResponse(httpBody []byte) ([]User, error) {
	userResponse := UserResponse{}

	err := json.Unmarshal(httpBody, &userResponse)

	if err != nil {
		return nil, err
	}

	return userResponse.Items, nil
}

func writeCSV(users []User) error {
	file, err := os.Create("users.csv")

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

func main() {
	log.Printf("Hitting up hipchat for all the users...")

	users, err := GetAllUsers()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d users!", len(users))

	err = writeCSV(users)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Wrote user data to users.csv")
}
