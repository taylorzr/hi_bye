package root

import(
	"encoding/json"
)

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	MentionName string `json:"mention_name"`
}

type Users []User

func (users *Users) Encode() ([]byte, error) {
	data, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (users *Users) Decode(data []byte) error {
	return json.Unmarshal(data, &users)
}
