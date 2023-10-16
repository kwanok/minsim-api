package comment

import (
	"context"
	"github.com/kwanok/minsim-api/database"
	"google.golang.org/grpc"
	"gorm.io/gorm/clause"
)

type commentService struct {
	CommentServer
	db database.DB
}

func (cs *commentService) NewComment(ctx context.Context, req *NewCommentRequest) (*NewCommentResponse, error) {
	comment := &database.RedditComment{
		ID:            req.Id,
		Author:        req.Author,
		Body:          req.Body,
		BodyHtml:      req.BodyHtml,
		CreatedUtc:    req.CreatedUtc,
		Distinguished: &req.Distinguished,
		Edited:        req.Edited,
		IsSubmitter:   req.IsSubmitter,
		LinkId:        req.LinkId,
		ParentId:      req.ParentId,
		Permalink:     req.Permalink,
		Stickied:      req.Stickied,
		Submission:    req.Submission,
		Subreddit:     req.Subreddit,
		SubredditId:   req.SubredditId,
	}

	err := cs.db.Session().Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"author":        comment.Author,
				"body":          comment.Body,
				"body_html":     comment.BodyHtml,
				"created_utc":   comment.CreatedUtc,
				"distinguished": comment.Distinguished,
				"edited":        comment.Edited,
				"is_submitter":  comment.IsSubmitter,
				"link_id":       comment.LinkId,
				"parent_id":     comment.ParentId,
				"permalink":     comment.Permalink,
				"stickied":      comment.Stickied,
				"submission":    comment.Submission,
				"subreddit":     comment.Subreddit,
				"subreddit_id":  comment.SubredditId,
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
