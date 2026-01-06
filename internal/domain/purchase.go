package domain

import "math"

// PurchaseMethod 購買方式
type PurchaseMethod string

const (
	MethodCard       PurchaseMethod = "card"       // 點卡儲值
	MethodCardReader PurchaseMethod = "cardreader" // 讀卡機
	MethodOriginal   PurchaseMethod = "original"   // 原價
	MethodGift       PurchaseMethod = "gift"       // 送禮
)

// CostPerDraw 每抽成本（點數）
const CostPerDraw = 27.0

// CalculatePoints 根據購買方式計算可得點數
func CalculatePoints(investment float64, method PurchaseMethod, discount float64) float64 {
	switch method {
	case MethodCard:
		// 點卡 xx 折：投入金額 ÷ 折數（四捨五入）
		if discount > 0 {
			return math.Round(investment / discount)
		}
		return investment
	case MethodCardReader:
		// 讀卡機 5% 回饋
		return investment * 1.05
	case MethodOriginal:
		// 原價
		return investment
	case MethodGift:
		// 送禮 xx 折：投入金額 ÷ 折數
		if discount > 0 {
			return math.Round(investment / discount)
		}
		return investment
	default:
		return investment
	}
}
