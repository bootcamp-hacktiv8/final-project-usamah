package entities

import (
	"database/sql"
	"time"
)

type Photo struct {
	Id         int          `json:"id"`
	User_id    int          `json:"user_id"`
	Title      string       `json:"title"`
	Caption    string       `json:"caption"`
	Photo_url  string       `json:"photo_url"`
	Created_at time.Time    `json:"created_at"`
	Updated_at sql.NullTime `json:"updated_at"`
	User       User         `json:"user"`
}
