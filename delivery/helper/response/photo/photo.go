package response

import (
	"database/sql"
	"final-project-usamah/entities"
	"time"
)

type PhotoFormatter struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	Created_at time.Time `json:"created_at"`
}

func FormatPhoto(photo entities.Photo) PhotoFormatter {
	formatter := PhotoFormatter{
		Id:         photo.Id,
		Title:      photo.Title,
		Caption:    photo.Caption,
		Photo_url:  photo.Photo_url,
		Created_at: photo.Created_at,
	}
	return formatter
}

type GetPhotoUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetPhotoFormatter struct {
	Id         int          `json:"id"`
	User_id    int          `json:"user_id"`
	Title      string       `json:"title"`
	Caption    string       `json:"caption"`
	Photo_url  string       `json:"photo_url"`
	Created_at time.Time    `json:"created_at"`
	Updated_at sql.NullTime `json:"updated_at"`
	User       GetPhotoUser `json:"user"`
}

type UpdatePhoto struct {
	Id         int          `json:"id"`
	User_id    int          `json:"user_id"`
	Title      string       `json:"title"`
	Caption    string       `json:"caption"`
	Photo_url  string       `json:"photo_url"`
	Updated_at sql.NullTime `json:"updated_at"`
}

func FormatUpdatePhoto(photo entities.Photo) UpdatePhoto {
	formatter := UpdatePhoto{
		Id:         photo.Id,
		User_id:    photo.User_id,
		Title:      photo.Title,
		Caption:    photo.Caption,
		Photo_url:  photo.Photo_url,
		Updated_at: photo.Updated_at,
	}
	return formatter
}
