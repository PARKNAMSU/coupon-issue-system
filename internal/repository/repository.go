package repository

import (
	"coupon-issuance-system.com/coupon-issuance-system/internal/model"
	"coupon-issuance-system.com/coupon-issuance-system/pkg/database"
	"github.com/jmoiron/sqlx"
)

type CouponRepository struct {
	database *sqlx.DB
}

var (
	repository *CouponRepository
)

func GetRepository() *CouponRepository {
	if repository == nil {
		repository = &CouponRepository{
			database: database.Connect(),
		}
	}
	return repository
}

func (r *CouponRepository) CreateCampagin(campagin model.CampaginEntity) (int, error) {
	result, err := r.database.Exec(`
	INSERT INTO campagin 
	SET name = ?,
	total_coupons = ?,
	avaliable_at = ?,
	status = ?
	`, campagin.Name, campagin.TotalCoupons, campagin.AvaliableAt, campagin.Status)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return int(id), err
}

func (r *CouponRepository) UpdateCampagin(campagin model.CampaginEntity) error {
	_, err := r.database.Exec(`
	UPDATE INTO campagin 
	SET name = ?,
	total_coupons = ?,
	avaliable_at = ?,
	status = ? 
	WHERE campagin_id = ? 
	`,
		campagin.Name,
		campagin.TotalCoupons,
		campagin.AvaliableAt,
		campagin.Status,
		campagin.CampaginId)

	return err
}

func (r *CouponRepository) GetCampaignById(campaginId int) (model.CampaginEntity, bool) {
	campaign := []model.CampaginEntity{}
	r.database.Select(
		&campaign,
		`SELECT * FROM campagin 
		WHERE campagin_id = ? 
		`,
		campaginId,
	)
	if len(campaign) == 0 {
		return model.CampaginEntity{}, false
	}
	return campaign[0], true
}

func (r *CouponRepository) IssueCoupon(coupon model.CouponEntity) error {
	_, err := r.database.Exec(`
	INSERT INTO coupon 
	SET campagin_id = ?,
	coupon_code = ?,
	is_used = 0
	`, coupon.CampaginId, coupon.CouponCode)
	return err
}

func (r *CouponRepository) GetCouponsByCampaignId(campaginId int) []model.CouponEntity {
	list := []model.CouponEntity{}
	r.database.Select(
		&list,
		`SELECT * FROM coupon 
		WHERE campagin_id = ? 
		`,
		campaginId,
	)
	return list
}
