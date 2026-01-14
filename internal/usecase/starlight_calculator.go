package usecase

import (
	"MSCashItemExpected/internal/domain"
	"math/rand"
	"time"
)

// StarlightCalculator 星光錦囊計算器
type StarlightCalculator struct {
	rng *rand.Rand
}

// NewStarlightCalculator 建立新的計算器
func NewStarlightCalculator() *StarlightCalculator {
	return &StarlightCalculator{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// CalculateEV 計算獎池的期望值
// 公式：EV = Σ(機率 × 價格)
func (sc *StarlightCalculator) CalculateEV(pool []domain.Reward) float64 {
	var ev float64
	for _, reward := range pool {
		ev += (reward.Probability / 100) * float64(reward.MarketPrice)
	}
	return ev
}

// CalculateContributions 計算各獎項對總期望值的貢獻
func (sc *StarlightCalculator) CalculateContributions(pool []domain.Reward) map[string]float64 {
	contributions := make(map[string]float64)
	for _, reward := range pool {
		contributions[reward.Name] = (reward.Probability / 100) * float64(reward.MarketPrice)
	}
	return contributions
}

// ExpandedItem 展開後的道具期望數量
type ExpandedItem struct {
	Name     string
	Expected float64 // 期望獲得數量（每抽）
}

// CalculateExpandedExpected 計算展開所有階段後的期望道具數量
// 會將「玲瓏星光」展開到星光結晶體、星光原石、星光水晶、璀璨星光
func (sc *StarlightCalculator) CalculateExpandedExpected(drawCount float64) []ExpandedItem {
	items := make(map[string]float64)

	// 輔助函數：累加道具
	addItem := func(name string, count float64) {
		items[name] += count
	}

	// 第一階段：星光錦囊
	for _, reward := range domain.Stage1Pool {
		if reward.Name == "玲瓏星光" {
			// 玲瓏星光需要展開
			// 4個玲瓏星光 = 1個星光結晶體
			crystalCount := drawCount * (reward.Probability / 100) / 4

			// 第二階段：星光結晶體
			for _, r2 := range domain.StagePools[2] {
				if r2.Name == "星光原石" {
					roughCount := crystalCount * (r2.Probability / 100)

					// 第三階段：星光原石
					for _, r3 := range domain.StagePools[3] {
						if r3.Name == "星光水晶" {
							pureCount := roughCount * (r3.Probability / 100)

							// 第四階段：星光水晶
							for _, r4 := range domain.StagePools[4] {
								if r4.Name == "璀璨星光" {
									brilliantCount := pureCount * (r4.Probability / 100)

									// 第五階段：璀璨星光（最終獎品）
									for _, r5 := range domain.StagePools[5] {
										addItem(r5.Name, brilliantCount*(r5.Probability/100))
									}
								} else {
									addItem(r4.Name, pureCount*(r4.Probability/100))
								}
							}
						} else {
							addItem(r3.Name, roughCount*(r3.Probability/100))
						}
					}
				} else {
					addItem(r2.Name, crystalCount*(r2.Probability/100))
				}
			}
		} else {
			addItem(reward.Name, drawCount*(reward.Probability/100))
		}
	}

	// 轉換為 slice
	var result []ExpandedItem
	for name, expected := range items {
		result = append(result, ExpandedItem{Name: name, Expected: expected})
	}

	return result
}

// SimulateStage1 第一階段模擬器
// 模擬大量開啟第一階段錦囊後的結果分佈
func (sc *StarlightCalculator) SimulateStage1(count int, pool []domain.Reward) domain.SimulationResult {
	results := make(map[string]int)
	crystalCount := 0

	for i := 0; i < count; i++ {
		reward := sc.drawFromPool(pool)
		results[reward.Name]++
		if reward.Name == "玲瓏星光" {
			crystalCount++
		}
	}

	// 計算理論期望的玲瓏星光數量
	theoreticalCrystal := float64(count) * (10.0 / 100) // 10% 機率

	// 計算總投入成本
	totalCost := count * domain.StarlightCost

	return domain.SimulationResult{
		DrawCount:     count,
		Results:       results,
		CrystalCount:  crystalCount,
		TheoreticalEV: theoreticalCrystal,
		TotalCost:     totalCost,
	}
}

// SimulateLadder 階梯升級模擬器
// 處理從第二階段到第五階段的邏輯
func (sc *StarlightCalculator) SimulateLadder(initialCount int) domain.LadderResult {
	result := domain.LadderResult{
		InitialCount: initialCount,
		Rewards:      make(map[string]int),
	}

	// 當前持有的錦囊數量（從第2階段開始）
	currentCount := initialCount

	// 第2階段：星光結晶體 -> 星光原石
	for i := 0; i < currentCount; i++ {
		reward := sc.drawFromPool(domain.StagePools[2])
		if reward.Name == domain.UpgradeItems[2] { // 星光原石
			// 升級成功，計入下一階段
			result.Rewards["星光原石"]++
		} else {
			// 失敗，記錄獎品並終止
			result.Stage2Failures++
			result.Rewards[reward.Name]++
		}
	}

	// 第3階段：星光原石 -> 星光水晶
	stage3Count := result.Rewards["星光原石"]
	delete(result.Rewards, "星光原石") // 原石已消耗
	for i := 0; i < stage3Count; i++ {
		reward := sc.drawFromPool(domain.StagePools[3])
		if reward.Name == domain.UpgradeItems[3] { // 星光水晶
			result.Rewards["星光水晶"]++
		} else {
			result.Stage3Failures++
			result.Rewards[reward.Name]++
		}
	}

	// 第4階段：星光水晶 -> 璀璨星光
	stage4Count := result.Rewards["星光水晶"]
	delete(result.Rewards, "星光水晶")
	for i := 0; i < stage4Count; i++ {
		reward := sc.drawFromPool(domain.StagePools[4])
		if reward.Name == domain.UpgradeItems[4] { // 璀璨星光
			result.Rewards["璀璨星光"]++
		} else {
			result.Stage4Failures++
			result.Rewards[reward.Name]++
		}
	}

	// 第5階段：璀璨星光 -> 最終獎品
	stage5Count := result.Rewards["璀璨星光"]
	delete(result.Rewards, "璀璨星光")
	result.Stage5Success = stage5Count

	for i := 0; i < stage5Count; i++ {
		reward := sc.drawFromPool(domain.StagePools[5])
		result.Rewards[reward.Name]++
	}

	return result
}

// drawFromPool 根據權重從獎池中抽取一個獎品
func (sc *StarlightCalculator) drawFromPool(pool []domain.Reward) domain.Reward {
	roll := sc.rng.Float64() * 100
	var cumulative float64

	for _, reward := range pool {
		cumulative += reward.Probability
		if roll < cumulative {
			return reward
		}
	}

	// 應該不會到這裡，但以防萬一返回最後一個
	return pool[len(pool)-1]
}

// CalculateSurvivalRate 計算存活率
func (sc *StarlightCalculator) CalculateSurvivalRate(result domain.LadderResult) float64 {
	if result.InitialCount == 0 {
		return 0
	}
	return float64(result.Stage5Success) / float64(result.InitialCount) * 100
}

// CalculateTheoreticalSurvival 計算理論存活率
// 每階段 50% 升級成功，需要通過 3 個升級檢查點
// 第2層→第3層(50%) → 第3層→第4層(50%) → 第4層→第5層(50%)
func (sc *StarlightCalculator) CalculateTheoreticalSurvival() float64 {
	// 0.5 ^ 3 = 0.125 = 12.5%
	return 0.5 * 0.5 * 0.5 * 100
}

// CalculateExpandedEV 計算展開後的期望總價值
func (sc *StarlightCalculator) CalculateExpandedEV(drawCount float64, prices map[string]int) float64 {
	items := sc.CalculateExpandedExpected(drawCount)
	var totalEV float64
	for _, item := range items {
		if price, ok := prices[item.Name]; ok {
			totalEV += item.Expected * float64(price)
		}
	}
	return totalEV
}

// IsZeroValueItem 檢查是否為價值為0的道具
func (sc *StarlightCalculator) IsZeroValueItem(name string) bool {
	for _, item := range domain.ZeroValueItems {
		if item == name {
			return true
		}
	}
	return false
}
