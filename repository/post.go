package repository

import (
	"database/sql"

	"github.com/wafi04/postservice/models"
)

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		DB: db,
	}
}

func (repo *PostRepository) Create(data models.Post) (*models.Post, error) {
	query := `
		INSERT INTO posts 
		(
			id,
			username,
			mentions,
			content,
			media,
			visibility,
			created_at,
			updated_at,
			parent_tweet_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING *
	`

	var post models.Post
	err := repo.DB.QueryRow(
		query,
		data.ID,
		data.Username,
		data.Mentions,
		data.Content,
		data.Media,
		data.Visibility,
		data.CreatedAt,
		data.UpdatedAt,
		data.ParentTweetID,
	).Scan(
		&post.ID,
		&post.Username,
		&post.Mentions,
		&post.Content,
		&post.Media,
		&post.Visibility,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.ParentTweetID,
	)

	if err != nil {
		return nil, err
	}

	return &post, nil
}
