package database

import "time"

type RedditSubmission struct {
	ID                string `gorm:"primaryKey"`
	Author            string
	AuthorFlairText   *string
	Clicked           bool
	CreatedUtc        int64
	IsOriginalContent bool
	IsSelf            bool
	LinkFlairText     *string
	Locked            bool
	Name              string
	NumComments       int64
	Over18            bool
	Permalink         string
	Score             int64
	Selftext          string
	Subreddit         string
	SubredditId       string
	Title             string
	UpvoteRatio       float64
	Url               string
}

type Minsim struct {
	ID        string `gorm:"primaryKey"`
	Type      string
	Positive  float64
	Negative  float64
	Neutral   float64
	Content   string
	Url       string
	User      string
	CreatedAt time.Time
}

type RedditComment struct {
	ID            string `gorm:"primaryKey"`
	Author        string
	Body          string
	BodyHtml      string
	CreatedUtc    int64
	Distinguished *string
	Edited        bool
	IsSubmitter   bool
	LinkId        string
	ParentId      string
	Permalink     string
	Stickied      bool
	Submission    string
	Subreddit     string
	SubredditId   string
}
