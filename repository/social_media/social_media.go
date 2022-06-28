package social_media

import (
	"context"
	"database/sql"
	_entities "final-project-usamah/entities"
	"time"

	_ "github.com/lib/pq"
)

type SocmedRepository struct {
	database *sql.DB
}

func NewSocmedRepository(db *sql.DB) *SocmedRepository {
	return &SocmedRepository{
		database: db,
	}
}

func (sr *SocmedRepository) CreateSocmed(newSocmed _entities.Social_media) (_entities.Social_media, int, error) {
	query := "INSERT INTO social_medias (user_id, name, social_media_url, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	ctx := context.Background()
	var id int
	err := sr.database.QueryRowContext(ctx, query, newSocmed.User_id, newSocmed.Name, newSocmed.Social_media_url, newSocmed.Created_at).Scan(&id)
	if err != nil {
		return newSocmed, id, err
	}
	return newSocmed, id, nil
}

func (sr *SocmedRepository) GetAllSocmed() ([]_entities.Social_media, error) {
	query := `SELECT social_medias.id, social_medias.user_id, social_medias.name, social_medias.social_media_url, social_medias.created_at, social_medias.updated_at,
	users.id, users.email, users.username
	FROM social_medias
	JOIN users ON (social_medias.user_id = users.id)`

	ctx := context.Background()

	rows, err := sr.database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var socmeds []_entities.Social_media
	for rows.Next() {
		var socmed _entities.Social_media
		err := rows.Scan(&socmed.Id, &socmed.User_id, &socmed.Name, &socmed.Social_media_url, &socmed.Created_at, &socmed.Updated_at, &socmed.User.Id, &socmed.User.Email, &socmed.User.Username)
		if err != nil {
			return nil, err
		}
		socmeds = append(socmeds, socmed)
	}
	return socmeds, nil
}

func (sr *SocmedRepository) GetSocmed(idSocmed int) (_entities.Social_media, error) {
	query := `SELECT id, user_id, name, social_media_url, created_at, updated_at
	FROM social_medias WHERE id = $1`
	ctx := context.Background()
	var socmed _entities.Social_media

	err := sr.database.QueryRowContext(ctx, query, idSocmed).Scan(&socmed.Id, &socmed.User_id, &socmed.Name, &socmed.Social_media_url, &socmed.Created_at, &socmed.Updated_at)
	if err != nil {
		return socmed, err
	}
	return socmed, nil
}

func (sr *SocmedRepository) UpdateSocmed(updateSocmed _entities.Social_media, idSocmed int) (_entities.Social_media, error) {
	query := `UPDATE social_medias SET name = $1, social_media_url = $2, updated_at = $3
	WHERE id = $4`
	ctx := context.Background()

	_, err := sr.database.ExecContext(ctx, query, updateSocmed.Name, updateSocmed.Social_media_url, time.Now(), idSocmed)
	if err != nil {
		return updateSocmed, err
	}
	return updateSocmed, nil
}

func (sr *SocmedRepository) DeleteSocmed(idSocmed int) error {
	query := `DELETE FROM social_medias WHERE id = $1`
	ctx := context.Background()

	_, err := sr.database.ExecContext(ctx, query, idSocmed)
	if err != nil {
		return err
	}
	return nil
}
