package usecase

import (
	"errors"
	"fmt"
	"time"

	"coupon-issuance-system.com/coupon-issuance-system/internal/model"
	"coupon-issuance-system.com/coupon-issuance-system/internal/repository"
	pb "coupon-issuance-system.com/coupon-issuance-system/proto"
)

type CouponUseCase struct {
	repository *repository.CouponRepository
}

var (
	usecase *CouponUseCase
)

func GetUseCase() *CouponUseCase {
	if usecase == nil {
		usecase = &CouponUseCase{
			repository: repository.GetRepository(),
		}
	}
	return usecase
}

func (u *CouponUseCase) CreateCampagin(input *pb.CreateCampaignRequest) (*pb.CreateCampaignResponse, error) {
	avaliableAt, err := time.Parse(time.RFC3339, input.AvaliableAt)
	if err != nil {
		return nil, err
	}

	id, err := u.repository.CreateCampagin(model.CampaginEntity{
		Name:         input.Name,
		AvaliableAt:  avaliableAt,
		TotalCoupons: int(input.TotalCoupons),
		Status:       int(input.Status),
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateCampaignResponse{
		CampaignId: int32(id),
	}, nil
}

func (u *CouponUseCase) UpdateCampagin(input *pb.UpdateCampaignRequest) (*pb.UpdateCampaignResponse, error) {
	avaliableAt, err := time.Parse(time.RFC3339, input.AvaliableAt)
	if err != nil {
		return nil, err
	}

	err = u.repository.UpdateCampagin(model.CampaginEntity{
		CampaginId:   int(input.CampaignId),
		Name:         input.Name,
		AvaliableAt:  avaliableAt,
		TotalCoupons: int(input.TotalCoupons),
		Status:       int(input.Status),
	})

	if err != nil {
		return nil, err
	}

	return &pb.UpdateCampaignResponse{
		Message: "updated",
	}, nil
}

func (u *CouponUseCase) GetCampagin(id int) (*pb.GetCampaignResponse, error) {
	campaign, isExist := u.repository.GetCampaignById(id)
	if !isExist {
		return nil, fmt.Errorf("not exist campaign id: %d", id)
	}

	returnData := &pb.GetCampaignResponse{
		CampaignId:   int32(id),
		Name:         campaign.Name,
		TotalCoupons: int32(campaign.TotalCoupons),
		AvaliableAt:  campaign.AvaliableAt.Format(time.RFC3339),
		Status:       int32(campaign.Status),
	}

	list := u.repository.GetCouponsByCampaignId(id)
	returnData.IssuedCoupons = make([]*pb.IssuedCoupon, 0, len(list))

	for _, data := range list {
		returnData.IssuedCoupons = append(returnData.IssuedCoupons, &pb.IssuedCoupon{
			CampaignId: int32(data.CampaginId),
			Code:       data.CouponCode,
		})
	}

	return returnData, nil
}

func (u *CouponUseCase) ValidationCampaign(data *pb.GetCampaignResponse) bool {
	avaliableAt, _ := time.Parse(time.RFC3339, data.AvaliableAt)
	if time.Now().Before(avaliableAt) {
		return false
	}
	return int(data.TotalCoupons) > len(data.IssuedCoupons)
}

func (u *CouponUseCase) IssueCoupon(input *pb.IssueCouponRequest) (*pb.IssueCouponResponse, error) {
	// todo 쿠폰 발급 로직 구현
	campaign, err := u.GetCampagin(int(input.CampaignId))

	if err != nil {
		return nil, err
	}

	if !u.ValidationCampaign(campaign) {
		return nil, errors.New("max coupon")
	}

	return nil, nil
}
