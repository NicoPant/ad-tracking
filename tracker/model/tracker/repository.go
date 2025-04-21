package tracker

import (
	"context"
	"github.com/google/uuid"
	"tracker/config"
	"tracker/db"
)

const Collection = "trackers"

type TrackerRepository interface {
	CreateTracker(ctx context.Context, adId string) (*Tracker, error)
	UpdateCountTracker(ctx context.Context, adId string) (*Tracker, error)
}

type TrackerService struct {
	cfg       *config.Config
	TrackerDb *TrackerRepository
}

func NewTrackerService(cfg *config.Config) *TrackerService {
	return &TrackerService{
		cfg: cfg,
	}
}

func (t *TrackerService) CreateTracker(ctx context.Context, adId string) (*Tracker, error) {
	collection := db.GetCollection(Collection, t.cfg)
	ctx, cancel := context.WithTimeout(ctx, t.cfg.Timeout)
	defer cancel()

	tracker := &Tracker{
		Id:    uuid.NewString(),
		AdId:  adId,
		Count: 0,
	}

	_, err := collection.InsertOne(ctx, tracker)
	if err != nil {
		return nil, err
	}

	return tracker, nil
}

func (t *TrackerService) UpdateCountTracker(ctx context.Context, adId string) (*Tracker, error) {
	collection := db.GetCollection(Collection, t.cfg)
	ctx, cancel := context.WithTimeout(ctx, t.cfg.Timeout)
	defer cancel()

	var tracker Tracker
	err := collection.FindOne(ctx, map[string]interface{}{"adId": adId}).Decode(&tracker)
	if err != nil {
		return nil, err
	}

	tracker.Count += 1

	_, err = collection.UpdateOne(ctx, map[string]interface{}{"adId": adId}, tracker)
	if err != nil {
		return nil, err
	}

	return &tracker, nil
}
