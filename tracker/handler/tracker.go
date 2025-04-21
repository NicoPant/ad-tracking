package handler

import (
	"context"
	"tracker/model/tracker"
)

type TrackerServiceServer struct {
	TrackerRepository tracker.TrackerRepository
}

func (h *TrackerServiceServer) CreateTracker(ctx context.Context, adId string) (*tracker.Tracker, error) {
	newTracker, err := h.TrackerRepository.CreateTracker(ctx, adId)
	if err != nil {
		return nil, err
	}
	return newTracker, nil
}

func (h *TrackerServiceServer) UpdateCountTracker(ctx context.Context, adId string) (*tracker.Tracker, error) {
	updatedTracker, err := h.TrackerRepository.UpdateCountTracker(ctx, adId)
	if err != nil {
		return nil, err
	}
	return updatedTracker, nil
}
