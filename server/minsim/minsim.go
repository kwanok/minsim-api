package minsim

import (
	"context"
	"github.com/kwanok/minsim-api/database"
	"google.golang.org/grpc"
	"gorm.io/gorm/clause"
	"time"
)

type minsimService struct {
	MinsimServer
	db database.DB
}

func (ms *minsimService) NewMinsim(ctx context.Context, req *NewMinsimRequest) (*NewMinsimResponse, error) {
	minsim := &database.Minsim{
		ID:        req.Id,
		Type:      req.Type,
		Positive:  float64(req.Positive),
		Negative:  float64(req.Negative),
		Neutral:   float64(req.Neutral),
		Content:   req.Content,
		Url:       req.Url,
		User:      req.User,
		CreatedAt: time.Unix(req.CreatedAt, 0),
	}

	err := ms.db.Session().Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"type":       req.Type,
				"positive":   req.Positive,
				"negative":   req.Negative,
				"neutral":    req.Neutral,
				"content":    req.Content,
				"url":        req.Url,
				"user":       req.User,
				"created_at": time.Unix(req.CreatedAt, 0),
			}),
		},
	).Create(&minsim).Error

	if err != nil {
		return nil, err
	}

	return &NewMinsimResponse{Id: minsim.ID}, nil
}

func Register(
	grpcServer *grpc.Server,
	db database.DB,
) {
	RegisterMinsimServer(grpcServer, &minsimService{db: db})
}
