package comment

import (
	"context"
	"github.com/kwanok/minsim-api/database"
	"google.golang.org/grpc"
	"gorm.io/gorm/clause"
	"time"
)

type commentService struct {
	CommentServer
	db database.DB
}

func (cs *commentService) NewComment(ctx context.Context, req *NewCommentRequest) (*NewCommentResponse, error) {
	comment := &database.RedditComment{
		ID:          req.Id,
		Content:     req.Content,
		Subreddit:   req.Subreddit,
		CreatedAt:   time.Unix(req.CreatedAt, 0),
		Author:      req.Author,
		SubredditId: req.SubredditId,
	}

	err := cs.db.Session().Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"content":      req.Content,
				"subreddit":    req.Subreddit,
				"created_at":   time.Unix(req.CreatedAt, 0),
				"author":       req.Author,
				"subreddit_id": req.SubredditId,
			}),
		},
	).Create(&comment).Error

	if err != nil {
		return nil, err
	}

	return &NewCommentResponse{Id: comment.ID}, nil
}

func Register(
	grpcServer *grpc.Server,
	db database.DB,
) {
	RegisterCommentServer(grpcServer, &commentService{db: db})
}
