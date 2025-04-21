package handler

import (
	"ad/model/ad"
	"ad/proto"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type AdServiceServer struct {
	proto.UnimplementedAdServiceServer
	AdRepository ad.AdRepository
}

func (h *AdServiceServer) CreateAd(ctx context.Context, req *proto.CreateAdRequest) (*proto.CreateAdResponse, error) {
	fmt.Println("CreateAd called", req)
	res, err := h.AdRepository.CreateAd(ctx, &ad.Ad{
		Id:          uuid.NewString(),
		Title:       req.Title,
		Description: req.Description,
		Url:         req.Url,
	})
	if err != nil {
		fmt.Println("Error creating ad:", err)
		return nil, err
	}
	fmt.Println("Ad created successfully:", res)

	return &proto.CreateAdResponse{
		Id: "12345",
	}, nil
}

func (h *AdServiceServer) GetAdById(ctx context.Context, req *proto.GetAdByIdRequest) (*proto.GetAdByIdResponse, error) {
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
