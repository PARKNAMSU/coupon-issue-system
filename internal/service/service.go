package service

import (
	"context"

	"coupon-issuance-system.com/coupon-issuance-system/internal/usecase"
	pb "coupon-issuance-system.com/coupon-issuance-system/proto"
)

type CouponService struct {
	pb.UnimplementedCouponServiceServer
	usecase *usecase.CouponUseCase
}

var (
	service *CouponService
)

func GetService() *CouponService {
	if service == nil {
		service = &CouponService{
			usecase: usecase.GetUseCase(),
		}
	}
	return service
}

func (s *CouponService) CreateCampagin(ctx context.Context, req *pb.CreateCampaignRequest) (*pb.CreateCampaignResponse, error) {
	return nil, nil
}

func (s *CouponService) UpdateCampaign(ctx context.Context, req *pb.UpdateCampaignRequest) (*pb.UpdateCampaignResponse, error) {
	return nil, nil
}

func (s *CouponService) GetCampaign(ctx context.Context, req *pb.GetCampaignRequest) (*pb.GetCampaignResponse, error) {
	return nil, nil
}

func (s *CouponService) IssueCoupon(ctx context.Context, req *pb.IssueCouponRequest) (*pb.IssueCouponResponse, error) {
	return nil, nil
}
