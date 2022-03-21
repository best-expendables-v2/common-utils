package trace

import "time"

type User struct {
	Id            string    `json:"id"`
	Email         string    `json:"email"`
	Active        bool      `json:"active"`
	Roles         []string  `json:"roles"`
	UpdatedAt     time.Time `json:"updatedAt"`
	CreatedAt     time.Time `json:"createdAt"`
}
