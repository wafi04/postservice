package models

import "time"

type Post struct {
	ID            string    `json:"id" db:"id"`
	Username      string    `json:"username" db:"username"`
	Mentions      string    `json:"mentions" db:"mentions"`
	Content       string    `json:"content" db:"content"`
	Media         *string   `json:"media,omitempty" db:"media"`
	LikesCount    int       `json:"likes_count" db:"likes_count"`
	RetweetsCount int       `json:"retweets_count" db:"retweets_count"`
	Visibility    string    `json:"visibility" db:"visibility"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	ParentTweetID *string   `json:"parent_tweet_id" db:"parent_tweet_id"`
}

type CreatePost struct {
	Username      string    `json:"username" db:"username"`
	Mentions      string    `json:"mentions" db:"mentions"`
	Content       string    `json:"content" db:"content"`
	Media         *string   `json:"media,omitempty" db:"media"`
	LikesCount    int       `json:"likes_count" db:"likes_count"`
	RetweetsCount int       `json:"retweets_count" db:"retweets_count"`
	Visibility    string    `json:"visibility" db:"visibility"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	ParentTweetID *string   `json:"parent_tweet_id" db:"parent_tweet_id"`
}
