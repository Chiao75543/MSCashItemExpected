package main

import (
	"MSCashItemExpected/internal/domain"
	"MSCashItemExpected/internal/usecase"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 需要輸入價值的道具（第一階段有價值道具）
var valuableItems = []string{
	"傳說潛在能力卷軸50%",
	"傳說潛在能力卷軸100%",
	"星力14星強化券",
	"星力15星強化券",
	"星力16星強化券",
	"星力17星強化券",
	"星力18星強化券",
	"星力19星強化券",
	"星力20星強化券",
	"突破1星強化券100%(21星)",
	"突破1星強化券100%(22星)",
	"追加1星強化券30%(23星)",
}

func main() {
	calculator := usecase.NewStarlightCalculator()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║          新楓之谷 星光錦囊 期望值計算器 & 模擬器              ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// ===========================================
	// 第一部分：輸入投入金額
	// ===========================================
	printSection("投入金額設定")

	fmt.Print("請輸入投入金額（台幣）: ")
	investmentStr, _ := reader.ReadString('\n')
	investment, _ := strconv.ParseFloat(strings.TrimSpace(investmentStr), 64)
	if investment <= 0 {
		investment = 10000 // 預設值
		fmt.Printf("使用預設值: %.0f 元\n", investment)
	}

	// 計算抽數（假設原價購買）
	drawCount := investment / float64(domain.StarlightCost)
	fmt.Printf("\n💰 投入金額: %.0f 元\n", investment)
	fmt.Printf("🎰 預計抽數: %.2f 次\n", drawCount)
	fmt.Println()

	// ===========================================
	// 第二部分：輸入道具價值
	// ===========================================
	printSection("道具價值設定（台幣）")

	fmt.Println("請輸入各道具的市場價值，直接按 Enter 表示價值為 0")
	fmt.Println()

	prices := make(map[string]int)

	for _, item := range valuableItems {
		fmt.Printf("  %s: ", item)
		priceStr, _ := reader.ReadString('\n')
		priceStr = strings.TrimSpace(priceStr)
		if priceStr != "" {
			if price, err := strconv.Atoi(priceStr); err == nil && price > 0 {
				prices[item] = price
			}
		}
	}

	// ===========================================
	// 第三部分：期望值計算結果
	// ===========================================
	printSection("期望值計算結果")

	// 建立機率對照表
	rateMap := make(map[string]float64)
	for _, r := range domain.Stage1Pool {
		rateMap[r.Name] = r.Probability
	}

	fmt.Println("┌────────────────────────────────┬─────────┬──────────┬────────────┐")
	fmt.Println("│ 道具名稱                       │  機率   │  單價    │  期望價值  │")
	fmt.Println("├────────────────────────────────┼─────────┼──────────┼────────────┤")

	var totalEV float64
	for _, item := range valuableItems {
		rate := rateMap[item]
		price := prices[item]
		expected := drawCount * (rate / 100)
		itemEV := expected * float64(price)
		totalEV += itemEV

		priceStr := "-"
		evStr := "-"
		if price > 0 {
			priceStr = fmt.Sprintf("%d", price)
			evStr = fmt.Sprintf("%.2f", itemEV)
		}

		fmt.Printf("│ %-30s │ %6.2f%% │ %8s │ %10s │\n",
			truncateName(item, 30),
			rate,
			priceStr,
			evStr)
	}

	fmt.Println("├────────────────────────────────┼─────────┼──────────┼────────────┤")
	fmt.Printf("│ %-30s │   ---   │   ---    │ %10.2f │\n", "【期望總價值】", totalEV)
	fmt.Println("└────────────────────────────────┴─────────┴──────────┴────────────┘")
	fmt.Println()

	// 報酬率計算
	roi := (totalEV - investment) / investment * 100
	roiSign := "+"
	if roi < 0 {
		roiSign = ""
	}

	fmt.Println("【投資報酬分析】")
	fmt.Printf("  投入金額: %.0f 元\n", investment)
	fmt.Printf("  期望回收: %.2f 元\n", totalEV)
	fmt.Printf("  期望報酬率: %s%.2f%%\n", roiSign, roi)
	fmt.Println()

	// ===========================================
	// 第四部分：模擬器
	// ===========================================
	printSection("是否執行模擬器？")

	fmt.Print("執行模擬器？(y/n): ")
	runSim, _ := reader.ReadString('\n')
	if strings.ToLower(strings.TrimSpace(runSim)) != "y" {
		fmt.Println("\n感謝使用！")
		return
	}

	// 第一階段模擬
	printSection("第一階段模擬器（模擬 1000 次開啟）")

	simResult := calculator.SimulateStage1(1000, domain.Stage1Pool)

	// 按數量排序
	type itemCount struct {
		name  string
		count int
	}
	var sorted []itemCount
	for name, count := range simResult.Results {
		sorted = append(sorted, itemCount{name, count})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].count > sorted[j].count
	})

	fmt.Println("┌────────────────────────────────┬──────────┬─────────────┐")
	fmt.Println("│ 道具名稱                       │   數量   │  佔比       │")
	fmt.Println("├────────────────────────────────┼──────────┼─────────────┤")

	for _, item := range sorted {
		percentage := float64(item.count) / float64(simResult.DrawCount) * 100
		fmt.Printf("│ %-30s │ %8d │ %10.2f%% │\n",
			truncateName(item.name, 30),
			item.count,
			percentage)
	}
	fmt.Println("└────────────────────────────────┴──────────┴─────────────┘")
	fmt.Println()

	// 玲瓏星光分析
	theoreticalCrystal := simResult.TheoreticalEV
	actualCrystal := float64(simResult.CrystalCount)
	deviation := actualCrystal - theoreticalCrystal
	deviationPct := (deviation / theoreticalCrystal) * 100

	fmt.Println("【玲瓏星光分析】")
	fmt.Printf("  理論期望數量: %.2f 個\n", theoreticalCrystal)
	fmt.Printf("  實際獲得數量: %d 個\n", simResult.CrystalCount)
	fmt.Printf("  偏差: %+.2f (%.2f%%)\n", deviation, deviationPct)
	fmt.Println()

	// 階梯升級模擬器
	printSection("大量階梯模擬（1000 個星光結晶體）")

	largeLadderResult := calculator.SimulateLadder(1000)
	printLadderResult(calculator, largeLadderResult)
}

func printSection(title string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 64))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 64))
	fmt.Println()
}

func truncateName(name string, maxLen int) string {
	runes := []rune(name)
	if len(runes) <= maxLen {
		return name + strings.Repeat(" ", maxLen-len(runes))
	}
	return string(runes[:maxLen-3]) + "..."
}

func printLadderResult(calculator *usecase.StarlightCalculator, result domain.LadderResult) {
	fmt.Println("【階段存活報告】")
	fmt.Println("┌─────────────────────┬──────────┬──────────┬──────────────┐")
	fmt.Println("│ 階段                │   進入   │   失敗   │   存活率     │")
	fmt.Println("├─────────────────────┼──────────┼──────────┼──────────────┤")

	stage2Survived := result.InitialCount - result.Stage2Failures
	stage3Survived := stage2Survived - result.Stage3Failures
	stage4Survived := stage3Survived - result.Stage4Failures

	survivalRate2 := float64(stage2Survived) / float64(result.InitialCount) * 100
	survivalRate3 := float64(stage3Survived) / float64(result.InitialCount) * 100
	survivalRate4 := float64(stage4Survived) / float64(result.InitialCount) * 100
	survivalRate5 := float64(result.Stage5Success) / float64(result.InitialCount) * 100

	fmt.Printf("│ 第2層（星光結晶體） │ %8d │ %8d │ %11.2f%% │\n",
		result.InitialCount, result.Stage2Failures, survivalRate2)
	fmt.Printf("│ 第3層（星光原石）   │ %8d │ %8d │ %11.2f%% │\n",
		stage2Survived, result.Stage3Failures, survivalRate3)
	fmt.Printf("│ 第4層（星光水晶）   │ %8d │ %8d │ %11.2f%% │\n",
		stage3Survived, result.Stage4Failures, survivalRate4)
	fmt.Printf("│ 第5層（璀璨星光）   │ %8d │    ---   │ %11.2f%% │\n",
		stage4Survived, survivalRate5)

	fmt.Println("└─────────────────────┴──────────┴──────────┴──────────────┘")
	fmt.Println()

	actualRate := calculator.CalculateSurvivalRate(result)
	theoreticalRate := calculator.CalculateTheoreticalSurvival()

	fmt.Println("【存活率分析】")
	fmt.Printf("  理論存活率: %.2f%% (0.5^3 = 12.5%%)\n", theoreticalRate)
	fmt.Printf("  實際存活率: %.2f%%\n", actualRate)
	fmt.Printf("  結論: 僅有 %.2f%% 的星光結晶體成功轉化為璀璨星光\n", actualRate)
	fmt.Println()

	fmt.Println("【獲得獎品統計】")
	fmt.Println("┌────────────────────────────────┬──────────┐")
	fmt.Println("│ 道具名稱                       │   數量   │")
	fmt.Println("├────────────────────────────────┼──────────┤")

	type itemCount struct {
		name  string
		count int
	}
	var sorted []itemCount
	for name, count := range result.Rewards {
		sorted = append(sorted, itemCount{name, count})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].count > sorted[j].count
	})

	for _, item := range sorted {
		fmt.Printf("│ %-30s │ %8d │\n",
			truncateName(item.name, 30),
			item.count)
	}
	fmt.Println("└────────────────────────────────┴──────────┘")
	fmt.Println()
}
