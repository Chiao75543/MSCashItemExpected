package main

import (
	"MSCashItemExpected/internal/adapter"
	"MSCashItemExpected/internal/usecase"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os/exec"
	"runtime"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	// 初始化各層（依賴注入）
	calculator := usecase.NewCalculator()
	handler := adapter.NewHandler(calculator)

	// 設定 API 路由
	http.HandleFunc("/api/calculate", handler.Calculate)

	// 設定靜態檔案服務
	staticFS, _ := fs.Sub(staticFiles, "static")
	http.Handle("/", http.FileServer(http.FS(staticFS)))

	port := "5278"
	url := fmt.Sprintf("http://localhost:%s", port)

	fmt.Println("=================================")
	fmt.Println("  新年氣息期望值計算機")
	fmt.Println("=================================")
	fmt.Printf("伺服器啟動於 %s\n", url)
	fmt.Println("按 Ctrl+C 結束程式")
	fmt.Println()

	// 自動開啟瀏覽器
	go openBrowser(url)

	// 啟動 HTTP 伺服器
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("伺服器啟動失敗: %v\n", err)
	}
}

// openBrowser 開啟預設瀏覽器
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	}

	if err != nil {
		fmt.Printf("無法自動開啟瀏覽器，請手動開啟: %s\n", url)
	}
}
