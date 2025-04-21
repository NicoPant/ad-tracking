package ad

import (
	"context"
	"fmt"
	"github.com/NicoPant/ad-tracking/ad/config"
	"github.com/NicoPant/ad-tracking/ad/db"
	"go.mongodb.org/mongo-driver/bson"
)

const Collection = "ads"

type AdRepository interface {
	GetAdById(ctx context.Context, id string) (*Ad, error)
	CreateAd(ctx context.Context, ad *Ad) (interface{}, error)
}

type AdService struct {
	cfg  *config.Config
	AdDb *AdRepository
}

func NewAdService(cfg *config.Config) *AdService {
	return &AdService{
		cfg: cfg,
	}
}

func (a *AdService) GetAdById(ctx context.Context, id string) (*Ad, error) {
	collection := db.GetCollection(Collection, a.cfg)
	ctx, cancel := context.WithTimeout(ctx, a.cfg.Timeout)
	defer cancel()

	var ad Ad
	err := collection.FindOne(ctx, bson.M{"_key": id}).Decode(&ad)
	if err != nil {
		return nil, err
	}
	fmt.Println(ad)
	return &ad, nil
}

func (a *AdService) CreateAd(ctx context.Context, ad *Ad) (interface{}, error) {
	fmt.Println("repository")
	collection := db.GetCollection(Collection, a.cfg)
	ctx, cancel := context.WithTimeout(ctx, a.cfg.Timeout)
	defer cancel()

	res, err := collection.InsertOne(ctx, ad)
	if err != nil {
		return nil, err
	}

	return res, nil
}
