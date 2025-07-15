package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/wafi04/postservice/models"
	"github.com/wafi04/postservice/repository"
)

type PostService struct {
	postRepo *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (service *PostService) Create(data models.CreatePost) (*models.Post, error) {
	var postId = uuid.New().String()
	return service.postRepo.Create(
		models.Post{
			ID:            postId,
			Username:      data.Username,
			Mentions:      data.Mentions,
			Content:       data.Content,
			Media:         data.Media,
			LikesCount:    0,
			RetweetsCount: 0,
			Visibility:    data.Visibility,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			ParentTweetID: data.ParentTweetID,
		},
	)
}
