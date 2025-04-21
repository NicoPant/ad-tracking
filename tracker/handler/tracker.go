package handler

import (
	"context"
	"github.com/NicoPant/ad-tracking/proto"
	"github.com/NicoPant/ad-tracking/tracker/model/tracker"
)

type TrackerServiceServer struct {
	proto.UnimplementedTrackerServiceServer
	TrackerRepository tracker.TrackerRepository
}

func (h *TrackerServiceServer) CreateTracker(ctx context.Context, request *proto.CreateTrackerRequest) (*proto.CreateTrackerResponse, error) {
	newTracker, err := h.TrackerRepository.CreateTracker(ctx, request.AdId)
	if err != nil {
		return nil, err
	}

	response := &proto.CreateTrackerResponse{
		Tracker: &proto.Tracker{
			Id:    newTracker.Id,
			AdId:  newTracker.AdId,
			Count: int32(newTracker.Count),
		},
	}
	return response, nil
}

func (h *TrackerServiceServer) UpdateCountTracker(ctx context.Context, request *proto.UpdateCountTrackerRequest) (*proto.UpdateCountTrackerResponse, error) {
	updatedTracker, err := h.TrackerRepository.UpdateCountTracker(ctx, request.AdId)
	if err != nil {
		return nil, err
	}

	response := &proto.UpdateCountTrackerResponse{
		Tracker: &proto.Tracker{
			Id:    updatedTracker.Id,
			AdId:  updatedTracker.AdId,
			Count: int32(updatedTracker.Count),
		},
	}
	return response, nil
}
