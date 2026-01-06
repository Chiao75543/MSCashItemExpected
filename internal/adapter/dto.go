package adapter

import (
	"MSCashItemExpected/internal/domain"
	"MSCashItemExpected/internal/usecase"
)

// CalculateRequest API 請求 DTO
type CalculateRequest struct {
	Investment float64   `json:"investment"`
	Method     string    `json:"method"`
	Discount   float64   `json:"discount"`
	BoxValues  BoxValues `json:"box_values"`
}

// BoxValues 心願箱價值 DTO
type BoxValues struct {
	Small  float64 `json:"small"`
	Medium float64 `json:"medium"`
	Large  float64 `json:"large"`
	Super  float64 `json:"super"`
}

// CalculateResponse API 回應 DTO
type CalculateResponse struct {
	Points          float64            `json:"points"`
	DrawCount       float64            `json:"draw_count"`
	CostPerBreath   float64            `json:"cost_per_breath"`
	ExpectedBreaths map[string]float64 `json:"expected_breaths"`
	ExpectedBoxes   map[string]float64 `json:"expected_boxes"`
	ExpectedValue   float64            `json:"expected_value"`
	ROI             float64            `json:"roi"`
}

// ToUseCaseInput 將 DTO 轉換為 UseCase 輸入
func (r CalculateRequest) ToUseCaseInput() usecase.CalculatorInput {
	return usecase.CalculatorInput{
		Investment: r.Investment,
		Method:     domain.PurchaseMethod(r.Method),
		Discount:   r.Discount,
		BoxValues: domain.BoxValues{
			Small:  r.BoxValues.Small,
			Medium: r.BoxValues.Medium,
			Large:  r.BoxValues.Large,
			Super:  r.BoxValues.Super,
		},
	}
}

// FromUseCaseOutput 將 UseCase 輸出轉換為 DTO
func FromUseCaseOutput(output usecase.CalculatorOutput) CalculateResponse {
	// 轉換 BreathCollection 為 map[string]float64
	breaths := make(map[string]float64)
	for zodiac, count := range output.ExpectedBreaths {
		breaths[string(zodiac)] = count
	}

	// 轉換 BoxCollection 為 map[string]float64
	boxes := make(map[string]float64)
	for boxType, count := range output.ExpectedBoxes {
		boxes[string(boxType)] = count
	}

	return CalculateResponse{
		Points:          output.Points,
		DrawCount:       output.DrawCount,
		CostPerBreath:   output.CostPerBreath,
		ExpectedBreaths: breaths,
		ExpectedBoxes:   boxes,
		ExpectedValue:   output.ExpectedValue,
		ROI:             output.ROI,
	}
}
