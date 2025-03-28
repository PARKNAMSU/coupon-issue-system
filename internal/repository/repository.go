package repository

import (
	"time"

	"coupon-issuance-system.com/coupon-issuance-system/internal/model"
	"coupon-issuance-system.com/coupon-issuance-system/pkg/database"
)

/*
	보완사항: 현재 인스턴스 메모리에 저장하는 데이터를 RDBMS 혹은 NoSQL 데이터베이스 저장소로 변경
*/

type CouponRepository struct {
	database *database.DataStore
}

var (
	repository *CouponRepository
)

func GetRepository() *CouponRepository {
	if repository == nil {
		repository = &CouponRepository{
			database: database.GetDataStore(),
		}
	}
	return repository
}

func (r *CouponRepository) CreateCampagin(campagin model.CampaginEntity) (int, error) {
	id := r.database.GetMaxCampaignsId()

	campagin.CreatedAt = time.Now()
	campagin.CampaginId = id + 1

	r.database.Campaigns[campagin.CampaginId] = campagin
	r.database.IncreaseMaxCampaignsId()

	return campagin.CampaginId, nil
}

func (r *CouponRepository) UpdateCampagin(campagin model.CampaginEntity) error {
	r.database.Campaigns[campagin.CampaginId] = campagin
	return nil
}

func (r *CouponRepository) GetCampaignById(campaginId int) (model.CampaginEntity, bool) {
	if data, isExist := r.database.Campaigns[campaginId]; !isExist {
		return model.CampaginEntity{}, false
	} else {
		return data, true
	}
}

func (r *CouponRepository) IssueCoupon(coupon model.CouponEntity) error {
	coupon.IsUsed = false
	coupon.CreatedAt = time.Now()
	r.database.Coupons[coupon.CampaginId][coupon.CouponCode] = coupon
	return nil
}

func (r *CouponRepository) GetCouponsByCampaignId(campaginId int) []model.CouponEntity {
	data, isExist := r.database.Coupons[campaginId]
	list := []model.CouponEntity{}

	if !isExist {
		return list
	}

	for _, v := range data {
		list = append(list, v)
	}

	return list
}
