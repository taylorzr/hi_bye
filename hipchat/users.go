package hipchat

import(
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/taylorzr/hibye/root"
)

type UserResponse struct {
	Items root.Users `json:"items"`
}

func GetAllUsers() (root.Users, error) {
	httpBody, err := getUsers()

	if err != nil {
		return nil, err
	}

	users, err := unmarshalUsers(httpBody)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func getUsers() ([]byte, error) {
	client := new(http.Client)

	url := fmt.Sprintf("https://api.hipchat.com/v2/user?max-results=1000&auth_token=%s", token)

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		message := fmt.Sprintf("Getting users: Expected 200 response, got %d", response.StatusCode)
		return nil, errors.New(message)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}


func unmarshalUsers(httpBody []byte) (root.Users, error) {
	userResponse := UserResponse{}

	err := json.Unmarshal(httpBody, &userResponse)

	if err != nil {
		return nil, err
	}

	return userResponse.Items, nil
}
