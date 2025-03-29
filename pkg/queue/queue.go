package queue

import (
	"fmt"
	"sync"

	"coupon-issuance-system.com/coupon-issuance-system/internal/model"
)

// 캠페인별 큐 구조체
type queue struct {
	ch chan model.CouponEntity // 쿠폰 요청을 처리할 채널
}

// IssueQueue 구조체
type IssueQueue struct {
	mu                sync.Mutex // 맵 접근을 보호하기 위한 Mutex
	queueByCampaignId map[int]chan model.CouponEntity
}

// IssueQueue 초기화 함수
func GetIssueQueue() *IssueQueue {
	return &IssueQueue{
		queueByCampaignId: make(map[int]chan model.CouponEntity),
	}
}

// 특정 캠페인에 쿠폰을 추가하는 함수
func (q *IssueQueue) InQueue(campaignId int, coupon model.CouponEntity) {
	q.mu.Lock()
	queueData, exists := q.queueByCampaignId[campaignId]

	if !exists {
		queueData = make(chan model.CouponEntity, 1)
		q.queueByCampaignId[campaignId] = queueData
	}
	q.mu.Unlock()

	// 채널 사용 시에는 Mutex를 사용하지 않음 (경쟁 조건 제거)
	queueData <- coupon
}

// 특정 캠페인의 쿠폰을 처리하는 함수
func (q *IssueQueue) DeQueue(campaignId int, issuFunc func(coupon model.CouponEntity)) error {
	q.mu.Lock()
	queueData, exists := q.queueByCampaignId[campaignId]
	q.mu.Unlock()

	if !exists {
		return fmt.Errorf("not exist task queue campaignId: %d", campaignId)
	}

	select {
	case coupon := <-queueData:
		issuFunc(coupon)
		return nil
	default:
		return fmt.Errorf("no coupons available in queue")
	}
}
