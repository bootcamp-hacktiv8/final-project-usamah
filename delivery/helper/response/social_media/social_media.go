package response

import (
	"database/sql"
	"final-project-usamah/entities"
	"time"
)

type SocialMediaFormatter struct {
	Id               int       `json:"id"`
	User_id          int       `json:"user_id"`
	Name             string    `json:"name"`
	Social_media_url string    `json:"social_media_url"`
	Created_at       time.Time `json:"created_at"`
}

func FormatSocialMedia(socialMedia entities.Social_media) SocialMediaFormatter {
	formatter := SocialMediaFormatter{
		Id:               socialMedia.Id,
		User_id:          socialMedia.User_id,
		Name:             socialMedia.Name,
		Social_media_url: socialMedia.Social_media_url,
		Created_at:       socialMedia.Created_at,
	}
	return formatter
}

type GetSosmedUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type FormatGetSosmed struct {
	Id               int           `json:"id"`
	User_id          int           `json:"user_id"`
	Name             string        `json:"name"`
	Social_media_url string        `json:"social_media_url"`
	Created_at       time.Time     `json:"created_at"`
	Updated_at       sql.NullTime  `json:"updated_at"`
	User             GetSosmedUser `json:"user"`
}

type UpdateSosmed struct {
	Id               int          `json:"id"`
	User_id          int          `json:"user_id"`
	Name             string       `json:"name"`
	Social_media_url string       `json:"social_media_url"`
	Updated_at       sql.NullTime `json:"updated_at"`
}

func FormatUpdateSosmed(sosmed entities.Social_media) UpdateSosmed {
	formatter := UpdateSosmed{
		Id:               sosmed.Id,
		User_id:          sosmed.User_id,
		Name:             sosmed.Name,
		Social_media_url: sosmed.Social_media_url,
		Updated_at:       sosmed.Updated_at,
	}
	return formatter
}
