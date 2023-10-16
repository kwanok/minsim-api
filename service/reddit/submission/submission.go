package submission

import (
	"context"
	"github.com/kwanok/minsim-api/database"
	"google.golang.org/grpc"
	"gorm.io/gorm/clause"
)

type submissionService struct {
	SubmissionServer
	db database.DB
}

func (ps *submissionService) NewSubmission(ctx context.Context, req *NewSubmissionRequest) (*NewSubmissionResponse, error) {
	redditPost := &database.RedditSubmission{
		ID:                req.Id,
		Author:            req.Author,
		AuthorFlairText:   &req.AuthorFlairText,
		Clicked:           req.Clicked,
		CreatedUtc:        req.CreatedUtc,
		IsOriginalContent: req.IsOriginalContent,
		IsSelf:            req.IsSelf,
		LinkFlairText:     &req.LinkFlairText,
		Locked:            req.Locked,
		Name:              req.Name,
		NumComments:       req.NumComments,
		Over18:            req.Over_18,
		Permalink:         req.Permalink,
		Score:             req.Score,
		Selftext:          req.Selftext,
		Subreddit:         req.Subreddit,
		SubredditId:       req.SubredditId,
		Title:             req.Title,
		UpvoteRatio:       float64(req.UpvoteRatio),
		Url:               req.Url,
	}

	err := ps.db.Session().Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"author":              redditPost.Author,
				"author_flair_text":   redditPost.AuthorFlairText,
				"clicked":             redditPost.Clicked,
				"created_utc":         redditPost.CreatedUtc,
				"is_original_content": redditPost.IsOriginalContent,
				"is_self":             redditPost.IsSelf,
				"link_flair_text":     redditPost.LinkFlairText,
				"locked":              redditPost.Locked,
				"name":                redditPost.Name,
				"num_comments":        redditPost.NumComments,
				"over18":              redditPost.Over18,
				"permalink":           redditPost.Permalink,
				"score":               redditPost.Score,
				"selftext":            redditPost.Selftext,
				"subreddit":           redditPost.Subreddit,
				"subreddit_id":        redditPost.SubredditId,
				"title":               redditPost.Title,
				"upvote_ratio":        redditPost.UpvoteRatio,
				"url":                 redditPost.Url,
			}),
		},
	).Create(&redditPost).Error

	if err != nil {
		return nil, err
	}

	return &NewSubmissionResponse{Id: redditPost.ID}, nil
}

func Register(
	grpcServer *grpc.Server,
	db database.DB,
) {
	RegisterSubmissionServer(grpcServer, &submissionService{db: db})
}
