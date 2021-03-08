package models

import (
	"encoding/json"
	"time"
)

// User data structure
type User struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	Weight    int16     `json:"weight"`
	Height    int16     `json:"height"`
}

// SampleUser ...
func SampleUser() User {
	user := User{
		ID:        1,
		Name:      "Mr. Bean",
		BirthDate: time.Date(1955, time.January, 6, 0, 0, 0, 0, time.UTC),
		Weight:    75,
		Height:    181,
	}
	return user
}

//UserToMap ...
func (u *User) UserToMap() (map[string]interface{}, error) {
	userMap := make(map[string]interface{})
	userJSON, err := json.Marshal(u)
	json.Unmarshal(userJSON, &userMap)

	return userMap, err
}
