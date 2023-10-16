package post

import (
	"context"
	"github.com/kwanok/minsim-api/database"
	"google.golang.org/grpc"
	"gorm.io/gorm/clause"
)

type postService struct {
	PostServer
	db database.DB
}

func (ps *postService) NewPost(ctx context.Context, req *NewPostRequest) (*NewPostResponse, error) {
	redditPost := &database.RedditPost{
		ID:        req.Id,
		Title:     req.Title,
		Url:       req.Url,
		Author:    req.Author,
		Content:   req.Content,
		Subreddit: req.Subreddit,
		Created:   req.CreatedAt,
	}

	err := ps.db.Session().Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"title":     req.Title,
				"author":    req.Author,
				"subreddit": req.Subreddit,
				"url":       req.Url,
				"content":   req.Content,
			}),
		},
	).Create(&redditPost).Error

	if err != nil {
		return nil, err
	}

	return &NewPostResponse{Id: redditPost.ID}, nil
}

func Register(
	grpcServer *grpc.Server,
	db database.DB,
) {
	RegisterPostServer(grpcServer, &postService{db: db})
}
