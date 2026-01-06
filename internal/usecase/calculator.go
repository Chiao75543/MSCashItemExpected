package usecase

import (
	"MSCashItemExpected/internal/domain"
)

// CalculatorInput 計算器輸入
type CalculatorInput struct {
	Investment float64
	Method     domain.PurchaseMethod
	Discount   float64
	BoxValues  domain.BoxValues
}

// CalculatorOutput 計算器輸出
type CalculatorOutput struct {
	Points          float64
	DrawCount       float64
	CostPerBreath   float64
	ExpectedBreaths domain.BreathCollection
	ExpectedBoxes   domain.BoxCollection
	ExpectedValue   float64
	ROI             float64
}

// Calculator 期望值計算器
type Calculator struct{}

// NewCalculator 建立計算器
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Calculate 計算期望值
func (c *Calculator) Calculate(input CalculatorInput) CalculatorOutput {
	// 1. 計算可得點數
	points := domain.CalculatePoints(input.Investment, input.Method, input.Discount)

	// 2. 計算可抽次數
	drawCount := points / domain.CostPerDraw

	// 3. 計算每個氣息的實際成本
	costPerBreath := 0.0
	if drawCount > 0 {
		costPerBreath = input.Investment / drawCount
	}

	// 4. 計算期望獲得各氣息數量
	expectedBreaths := c.calculateExpectedBreaths(drawCount)

	// 5. 計算期望可湊心願箱數量（貪心算法）
	expectedBoxes := c.calculateExpectedBoxes(expectedBreaths)

	// 6. 計算期望總價值
	expectedValue := expectedBoxes.TotalValue(input.BoxValues)

	// 7. 計算報酬率
	roi := 0.0
	if input.Investment > 0 {
		roi = ((expectedValue - input.Investment) / input.Investment) * 100
	}

	return CalculatorOutput{
		Points:          points,
		DrawCount:       drawCount,
		CostPerBreath:   costPerBreath,
		ExpectedBreaths: expectedBreaths,
		ExpectedBoxes:   expectedBoxes,
		ExpectedValue:   expectedValue,
		ROI:             roi,
	}
}

// calculateExpectedBreaths 計算期望獲得的氣息數量
func (c *Calculator) calculateExpectedBreaths(drawCount float64) domain.BreathCollection {
	breaths := domain.NewBreathCollection()
	for zodiac, rate := range domain.ZodiacRates {
		breaths[zodiac] = drawCount * (rate / 100.0)
	}
	return breaths
}

// calculateExpectedBoxes 計算期望心願箱數量（貪心算法，優先湊高價值）
func (c *Calculator) calculateExpectedBoxes(breaths domain.BreathCollection) domain.BoxCollection {
	// 複製一份，避免修改原始資料
	remaining := breaths.Clone()
	result := domain.NewBoxCollection()

	// 按優先順序湊箱
	for _, boxType := range domain.BoxPriority {
		requirements := domain.BoxRequirements[boxType]

		// 找出可湊的箱數（取最小值）
		minCount := remaining.Min(requirements)

		if minCount > 0 {
			result[boxType] = minCount
			// 扣除已使用的氣息
			remaining.Subtract(requirements, minCount)
		}
	}

	return result
}
