package entities

import (
	"database/sql"
	"time"
)

type Social_media struct {
	Id               int          `json:"id"`
	User_id          int          `json:"user_id"`
	Name             string       `json:"name"`
	Social_media_url string       `json:"social_media_url"`
	Created_at       time.Time    `json:"created_at"`
	Updated_at       sql.NullTime `json:"updated_at"`
}
