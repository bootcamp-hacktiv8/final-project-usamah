package entities

import (
	"database/sql"
	"time"
)

type Comment struct {
	Id         int          `json:"id"`
	User_id    int          `json:"user_id"`
	Photo_id   int          `json:"photo_id"`
	Message    string       `json:"message"`
	Created_at time.Time    `json:"created_at"`
	Updated_at sql.NullTime `json:"updated_at"`
	User       User         `json:"user"`
	Photo      Photo        `json:"photo"`
}
