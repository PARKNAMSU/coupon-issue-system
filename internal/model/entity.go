package model

import (
	"database/sql"
	"time"
)

type CampaginEntity struct { // campaginId UNIQ
	CampaginId   int          `db:"campagin_id"`
	Name         string       `db:"name"`          // 캠페인 이름
	AvaliableAt  time.Time    `db:"avaliable_at"`  // 사용 가능 날짜
	TotalCoupons int          `db:"total_coupons"` // 총 쿠폰 갯수
	Status       int          `db:"status"`        // 0: 닫힘, 1: 열림
	CreatedAt    time.Time    `db:"created_at"`    // 생성 날짜
	UpdatedAt    sql.NullTime `db:"updated_at"`    // 최근 업데이트 날짜
}

type CouponEntity struct { // (campaginId, couponCode) UNIQ
	CampaginId int          `db:"campagin_id"` // 연결된 캠페인
	CouponCode string       `db:"coupon_code"` // 쿠폰 코드
	IsUsed     bool         `db:"isUsed"`      // 쿠폰 사용 번호
	CreatedAt  time.Time    `db:"created_at"`  // 생성 날짜
	UsedAt     sql.NullTime `db:"used_at"`     // 사용날짜
}
