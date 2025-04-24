package handler

import (
	"context"
	"fmt"

	"github.com/NicoPant/ad-tracking/ad/model/ad"
	"github.com/NicoPant/ad-tracking/proto"
	"github.com/google/uuid"
)

type AdServiceServer struct {
	proto.UnimplementedAdServiceServer
	AdRepository  ad.AdRepository
	TrackerClient proto.TrackerServiceClient
}

func (h *AdServiceServer) CreateAd(ctx context.Context, req *proto.CreateAdRequest) (*proto.CreateAdResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}
	if req.Title == "" || req.Description == "" || req.Url == "" {
		return nil, fmt.Errorf("title, description, and url are required")
	}
	newUuid := uuid.NewString()
	res, err := h.AdRepository.CreateAd(ctx, &ad.Ad{
		Id:          newUuid,
		Title:       req.Title,
		Description: req.Description,
		Url:         req.Url,
	})
	if err != nil {
		fmt.Println("Error creating ad:", err)
		return nil, err
	}
	fmt.Println("Ad created successfully:", res)

	_, err = h.TrackerClient.CreateTracker(ctx, &proto.CreateTrackerRequest{AdId: newUuid})
	if err != nil {
		fmt.Println("Error creating tracker:", err)
		return nil, err
	}

	return &proto.CreateAdResponse{
		Id: newUuid,
	}, nil
}

func (h *AdServiceServer) GetAdById(ctx context.Context, req *proto.GetAdByIdRequest) (*proto.GetAdByIdResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}
	if req.Id == "" {
		return nil, fmt.Errorf("ad ID is required")
	}

	res, err := h.AdRepository.GetAdById(ctx, req.Id)
	if err != nil {
		fmt.Println("Error getting ad by ID:", err)
		return nil, err
	}
	fmt.Println("Ad retrieved successfully:", res)
	return &proto.GetAdByIdResponse{
		Ad: &proto.Ad{
			Id:          res.Id,
			Title:       res.Title,
			Description: res.Description,
			Url:         res.Url,
		},
	}, nil
}

func (h *AdServiceServer) ServeAd(ctx context.Context, req *proto.ServeAdRequest) (*proto.ServeAdResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}
	if req.AdId == "" {
		return nil, fmt.Errorf("ad ID is required")
	}
	res, err := h.AdRepository.GetAdById(ctx, req.AdId)
	if err != nil {
		fmt.Println("Error getting ad by ID:", err)
		return nil, err
	}
	fmt.Println("Ad retrieved successfully:", res)

	_, err = h.TrackerClient.UpdateCountTracker(ctx, &proto.UpdateCountTrackerRequest{AdId: req.AdId})
	if err != nil {
		fmt.Println("Error updating tracker count:", err)
		return nil, err
	}

	return &proto.ServeAdResponse{
		Url: res.Url,
	}, nil
}
