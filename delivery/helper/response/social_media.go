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

func ResponseSocialMedia(socialMedia entities.Social_media) SocialMediaFormatter {
	formatter := SocialMediaFormatter{
		Id:               socialMedia.Id,
		User_id:          socialMedia.User_id,
		Name:             socialMedia.Name,
		Social_media_url: socialMedia.Social_media_url,
		Created_at:       socialMedia.Created_at,
	}
	return formatter
}

type GetSocmedUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetSocmedFormatter struct {
	Id               int           `json:"id"`
	User_id          int           `json:"user_id"`
	Name             string        `json:"name"`
	Social_media_url string        `json:"social_media_url"`
	Created_at       time.Time     `json:"created_at"`
	Updated_at       sql.NullTime  `json:"updated_at"`
	User             GetSocmedUser `json:"user"`
}

func ResponseGetSocialMedia(socialMedia []entities.Social_media) []GetSocmedFormatter {
	var sosmeds []GetSocmedFormatter
	for i := 0; i < len(socialMedia); i++ {
		sosmed := GetSocmedFormatter{
			Id:               socialMedia[i].Id,
			User_id:          socialMedia[i].User_id,
			Name:             socialMedia[i].Name,
			Social_media_url: socialMedia[i].Social_media_url,
			Created_at:       socialMedia[i].Created_at,
			Updated_at:       socialMedia[i].Updated_at,
			User: GetSocmedUser{
				Id:       socialMedia[i].User.Id,
				Username: socialMedia[i].User.Username,
				Email:    socialMedia[i].User.Email,
			},
		}
		sosmeds = append(sosmeds, sosmed)
	}
	return sosmeds
}

type UpdateSocmedFormatter struct {
	Id               int          `json:"id"`
	User_id          int          `json:"user_id"`
	Name             string       `json:"name"`
	Social_media_url string       `json:"social_media_url"`
	Updated_at       sql.NullTime `json:"updated_at"`
}

func ResponseUpdateSocmed(sosmed entities.Social_media) UpdateSocmedFormatter {
	formatter := UpdateSocmedFormatter{
		Id:               sosmed.Id,
		User_id:          sosmed.User_id,
		Name:             sosmed.Name,
		Social_media_url: sosmed.Social_media_url,
		Updated_at:       sosmed.Updated_at,
	}
	return formatter
}
