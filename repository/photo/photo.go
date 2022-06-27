package photo

import (
	"context"
	"database/sql"
	_entities "final-project-usamah/entities"
	"time"

	_ "github.com/lib/pq"
)

type PhotoRepository struct {
	database *sql.DB
}

func NewPhotoRepository(db *sql.DB) *PhotoRepository {
	return &PhotoRepository{
		database: db,
	}
}

func (pr *PhotoRepository) CreatePhoto(newPhoto _entities.Photo) (_entities.Photo, int, error) {
	query := "INSERT INTO photos (user_id, title, caption, photo_url, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	ctx := context.Background()
	var id int
	err := pr.database.QueryRowContext(ctx, query, newPhoto.User_id, newPhoto.Title, newPhoto.Caption, newPhoto.Photo_url, time.Now()).Scan(&id)
	if err != nil {
		return newPhoto, id, err
	}
	return newPhoto, id, nil
}

func (pr *PhotoRepository) GetAllPhoto() ([]_entities.Photo, error) {
	query := `SELECT photos.id, photos.user_id, photos.title, photos.caption, photos.photo_url,
	photos.created_at, photos.updated_at, users.email, users.username
	FROM photos
	JOIN users ON (photos.user_id = users.id)`

	ctx := context.Background()

	rows, err := pr.database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []_entities.Photo
	for rows.Next() {
		var photo _entities.Photo
		err := rows.Scan(&photo.Id, &photo.User_id, &photo.Title, &photo.Caption, &photo.Photo_url, &photo.Created_at, &photo.Updated_at, &photo.User.Email, &photo.User.Username)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

func (pr *PhotoRepository) GetPhoto(idPhoto int) (_entities.Photo, error) {
	query := `SELECT id, user_id, title, caption, photo_url, created_at, updated_at
	FROM photos WHERE id = $1`
	ctx := context.Background()
	var photo _entities.Photo

	err := pr.database.QueryRowContext(ctx, query, idPhoto).Scan(&photo.Id, &photo.User_id, &photo.Title, &photo.Caption, &photo.Photo_url, &photo.Created_at, &photo.Updated_at)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (pr *PhotoRepository) UpdatePhoto(updatePhoto _entities.Photo, idPhoto int) (_entities.Photo, error) {
	query := `UPDATE photos SET title = $1, caption = $2, photo_url = $3, updated_at = $4
	WHERE id = $5`
	ctx := context.Background()

	_, err := pr.database.ExecContext(ctx, query, updatePhoto.Title, updatePhoto.Caption, updatePhoto.Photo_url, time.Now(), idPhoto)
	if err != nil {
		return updatePhoto, err
	}
	return updatePhoto, nil
}

func (pr *PhotoRepository) DeletePhoto(idPhoto int) error {
	query := `DELETE FROM photos WHERE id = $1`
	ctx := context.Background()

	_, err := pr.database.ExecContext(ctx, query, idPhoto)
	if err != nil {
		return err
	}
	return nil
}
