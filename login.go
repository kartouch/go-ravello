package ravello

import (
	"encoding/json"
)

//User are objects responses from login endpoint
type User struct {
	Actived        string   `json:"actived,omitempty"`
	Email          string   `json:"email,omitempty"`
	Enabled        bool     `json:"enabled,omitempty"`
	ID             uint64   `json:"id,omitempty"`
	InvitationTime uint64   `json:"invitationTime,omitempty"`
	Name           string   `json:"name,omitempty"`
	Surname        string   `json:"surname,omitempty"`
	Organization   uint64   `json:"organization,omitempty"`
	Roles          []string `json:"roles,omitempty"`
	UUID           string   `json:"uuid,omitempty"`
}

//Login returns user details
func Login() (u User, err error) {
	data, err := handler("POST", "/login", nil)
	json.Unmarshal(data, &u)
	return
}
