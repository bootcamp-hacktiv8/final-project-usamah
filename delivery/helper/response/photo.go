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

func ResponsePhoto(photo entities.Photo) PhotoFormatter {
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

func ResponseGetPhoto(photo []entities.Photo) []GetPhotoFormatter {
	var photos []GetPhotoFormatter
	for i := 0; i < len(photo); i++ {
		getphoto := GetPhotoFormatter{
			Id:         photo[i].Id,
			User_id:    photo[i].User_id,
			Title:      photo[i].Title,
			Caption:    photo[i].Caption,
			Photo_url:  photo[i].Photo_url,
			Created_at: photo[i].Created_at,
			Updated_at: photo[i].Updated_at,
			User: GetPhotoUser{
				Username: photo[i].User.Username,
				Email:    photo[i].User.Email,
			},
		}
		photos = append(photos, getphoto)
	}
	return photos
}

type UpdatePhotoFormatter struct {
	Id         int          `json:"id"`
	User_id    int          `json:"user_id"`
	Title      string       `json:"title"`
	Caption    string       `json:"caption"`
	Photo_url  string       `json:"photo_url"`
	Updated_at sql.NullTime `json:"updated_at"`
}

func ResponseUpdatePhoto(photo entities.Photo) UpdatePhotoFormatter {
	formatter := UpdatePhotoFormatter{
		Id:         photo.Id,
		User_id:    photo.User_id,
		Title:      photo.Title,
		Caption:    photo.Caption,
		Photo_url:  photo.Photo_url,
		Updated_at: photo.Updated_at,
	}
	return formatter
}
