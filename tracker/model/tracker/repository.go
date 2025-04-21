package tracker

import (
	"context"
	"fmt"
	"github.com/NicoPant/ad-tracking/tracker/config"
	"github.com/NicoPant/ad-tracking/tracker/db"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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
	err := collection.FindOne(ctx, map[string]interface{}{"ad_id": adId}).Decode(&tracker)
	if err != nil {
		fmt.Println("Error finding tracker:", err)
		return nil, err
	}

	tracker.Count += 1

	filter := bson.M{"ad_id": tracker.AdId}
	update := bson.M{
		"$set": bson.M{
			"count": tracker.Count,
		},
	}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &tracker, nil
}
