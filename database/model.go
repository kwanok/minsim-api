package database

import "time"

type RedditPost struct {
	ID        string `gorm:"primaryKey"`
	Subreddit string
	Author    string
	Title     string
	Url       string
	Content   string
	Created   int64
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
	ID          string `gorm:"primaryKey"`
	Author      string
	Subreddit   string
	CreatedAt   time.Time
	Content     string
	SubredditId string
}
