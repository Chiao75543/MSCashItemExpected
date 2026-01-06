package adapter

import (
	"MSCashItemExpected/internal/usecase"
	"encoding/json"
	"net/http"
)

// Handler HTTP 處理器
type Handler struct {
	calculator *usecase.Calculator
}

// NewHandler 建立 Handler
func NewHandler(calculator *usecase.Calculator) *Handler {
	return &Handler{
		calculator: calculator,
	}
}

// Calculate 處理計算請求
func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析請求
	var req CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 轉換為 UseCase 輸入
	input := req.ToUseCaseInput()

	// 執行計算
	output := h.calculator.Calculate(input)

	// 轉換為回應 DTO
	response := FromUseCaseOutput(output)

	// 回傳 JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
