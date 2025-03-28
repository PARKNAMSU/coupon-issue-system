package database

import "coupon-issuance-system.com/coupon-issuance-system/internal/model"

// 임시 데이터베이스 (해당 패키지를 데이터베이스라고 가정)

type DataStore struct {
	Campaigns      map[int]model.CampaginEntity // campaignId key
	maxCampaignsId int
	Coupons        map[int]map[string]model.CouponEntity // campaignId, code key
}

var (
	dataStore = &DataStore{
		Campaigns:      map[int]model.CampaginEntity{},
		Coupons:        map[int]map[string]model.CouponEntity{},
		maxCampaignsId: 0,
	}
)

func GetDataStore() *DataStore {
	return dataStore
}

func (d *DataStore) GetMaxCampaignsId() int {
	return d.maxCampaignsId
}

func (d *DataStore) IncreaseMaxCampaignsId() {
	d.maxCampaignsId += 1
}
