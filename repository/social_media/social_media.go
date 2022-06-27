package social_media

import (
	"context"
	"database/sql"
	response "final-project-usamah/delivery/helper/response/social_media"
	_entities "final-project-usamah/entities"
	"time"

	_ "github.com/lib/pq"
)

type SosmedRepository struct {
	database *sql.DB
}

func NewSosmedRepository(db *sql.DB) *SosmedRepository {
	return &SosmedRepository{
		database: db,
	}
}

func (sr *SosmedRepository) CreateSosmed(newSosmed _entities.Social_media) (_entities.Social_media, int, error) {
	query := "INSERT INTO social_medias (user_id, name, social_media_url, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	ctx := context.Background()
	var id int
	err := sr.database.QueryRowContext(ctx, query, newSosmed.User_id, newSosmed.Name, newSosmed.Social_media_url, newSosmed.Created_at).Scan(&id)
	if err != nil {
		return newSosmed, id, err
	}
	return newSosmed, id, nil
}

func (sr *SosmedRepository) GetAllSosmed() ([]response.FormatGetSosmed, error) {
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

	var sosmeds []response.FormatGetSosmed
	for rows.Next() {
		var sosmed response.FormatGetSosmed
		err := rows.Scan(&sosmed.Id, &sosmed.User_id, &sosmed.Name, &sosmed.Social_media_url, &sosmed.Created_at, &sosmed.Updated_at, &sosmed.User.Id, &sosmed.User.Email, &sosmed.User.Username)
		if err != nil {
			return nil, err
		}
		sosmeds = append(sosmeds, sosmed)
	}
	return sosmeds, nil
}

func (sr *SosmedRepository) GetSosmed(idSosmed int) (_entities.Social_media, error) {
	query := `SELECT id, user_id, name, social_media_url, created_at, updated_at
	FROM social_medias WHERE id = $1`
	ctx := context.Background()
	var sosmed _entities.Social_media

	err := sr.database.QueryRowContext(ctx, query, idSosmed).Scan(&sosmed.Id, &sosmed.User_id, &sosmed.Name, &sosmed.Social_media_url, &sosmed.Created_at, &sosmed.Updated_at)
	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

func (sr *SosmedRepository) UpdateSosmed(updateSosmed _entities.Social_media, idSosmed int) (_entities.Social_media, error) {
	query := `UPDATE social_medias SET name = 1$, social_media_url = $2, updated_at = $3
	WHERE id = $4`
	ctx := context.Background()

	_, err := sr.database.ExecContext(ctx, query, updateSosmed.Name, updateSosmed.Social_media_url, time.Now(), idSosmed)
	if err != nil {
		return updateSosmed, err
	}
	return updateSosmed, nil
}

func (sr *SosmedRepository) DeleteSosmed(idSosmed int) error {
	query := `DELETE FROM social_medias WHERE id = $1`
	ctx := context.Background()

	_, err := sr.database.ExecContext(ctx, query, idSosmed)
	if err != nil {
		return err
	}
	return nil
}
