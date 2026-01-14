package domain

// Reward 獎品結構
type Reward struct {
	Name        string  // 道具名稱
	Probability float64 // 機率 (%)
	MarketPrice int     // 市場價值
}

// StarlightCost 每抽成本（點數）
const StarlightCost = 45

// Stage1Pool 第一階段（星光錦囊）獎池
var Stage1Pool = []Reward{
	{Name: "靈魂艾爾達碎片交換券(10個)", Probability: 8.00, MarketPrice: 0},
	{Name: "靈魂艾爾達", Probability: 6.00, MarketPrice: 0},
	{Name: "永遠的輪迴星火", Probability: 14.40, MarketPrice: 0},
	{Name: "暗黑輪迴星火", Probability: 13.70, MarketPrice: 0},
	{Name: "特別附加潛在能力賦予卷軸", Probability: 7.80, MarketPrice: 0},
	{Name: "傳說潛在能力卷軸50%", Probability: 0.85, MarketPrice: 0},
	{Name: "傳說潛在能力卷軸100%", Probability: 0.55, MarketPrice: 0},
	{Name: "星力14星強化券", Probability: 15.00, MarketPrice: 0},
	{Name: "星力15星強化券", Probability: 10.00, MarketPrice: 0},
	{Name: "星力16星強化券", Probability: 7.00, MarketPrice: 0},
	{Name: "星力17星強化券", Probability: 3.40, MarketPrice: 0},
	{Name: "星力18星強化券", Probability: 1.50, MarketPrice: 0},
	{Name: "星力19星強化券", Probability: 0.60, MarketPrice: 0},
	{Name: "星力20星強化券", Probability: 0.40, MarketPrice: 0},
	{Name: "突破1星強化券100%(21星)", Probability: 0.45, MarketPrice: 0},
	{Name: "突破1星強化券100%(22星)", Probability: 0.20, MarketPrice: 0},
	{Name: "追加1星強化券30%(23星)", Probability: 0.15, MarketPrice: 0},
	{Name: "玲瓏星光", Probability: 10.00, MarketPrice: 0}, // 二階錦囊
}

// StagePools 階梯式獎池（第2-5階段）
var StagePools = map[int][]Reward{
	// 第2階段：星光結晶體（4個玲瓏星光合成）
	2: {
		{Name: "星力18星強化券", Probability: 18.00, MarketPrice: 0},
		{Name: "星力19星強化券", Probability: 12.00, MarketPrice: 0},
		{Name: "星力20星強化券", Probability: 6.00, MarketPrice: 0},
		{Name: "突破1星強化券30%(23星)", Probability: 10.00, MarketPrice: 0},
		{Name: "突破1星強化券50%(23星)", Probability: 4.00, MarketPrice: 0},
		{Name: "星光原石", Probability: 50.00, MarketPrice: 0}, // 升級成功
	},
	// 第3階段：星光原石
	3: {
		{Name: "星力19星強化券", Probability: 10.00, MarketPrice: 0},
		{Name: "星力20星強化券", Probability: 8.00, MarketPrice: 0},
		{Name: "星力21星強化券", Probability: 2.00, MarketPrice: 0},
		{Name: "突破1星強化券30%(23星)", Probability: 8.00, MarketPrice: 0},
		{Name: "突破1星強化券50%(23星)", Probability: 6.00, MarketPrice: 0},
		{Name: "突破1星強化券100%(23星)", Probability: 5.00, MarketPrice: 0},
		{Name: "突破1星強化券30%(24星)", Probability: 7.00, MarketPrice: 0},
		{Name: "突破1星強化券50%(24星)", Probability: 4.00, MarketPrice: 0},
		{Name: "星光水晶", Probability: 50.00, MarketPrice: 0}, // 升級成功
	},
	// 第4階段：星光水晶
	4: {
		{Name: "突破1星強化券50%(23星)", Probability: 20.00, MarketPrice: 0},
		{Name: "突破1星強化券100%(23星)", Probability: 15.00, MarketPrice: 0},
		{Name: "突破1星強化券30%(24星)", Probability: 8.00, MarketPrice: 0},
		{Name: "突破1星強化券50%(24星)", Probability: 4.00, MarketPrice: 0},
		{Name: "突破1星強化券100%(24星)", Probability: 2.00, MarketPrice: 0},
		{Name: "突破1星強化券30%(25星)", Probability: 0.70, MarketPrice: 0},
		{Name: "突破1星強化券50%(25星)", Probability: 0.30, MarketPrice: 0},
		{Name: "璀璨星光", Probability: 50.00, MarketPrice: 0}, // 升級成功
	},
	// 第5階段：璀璨星光（最終階段）
	5: {
		{Name: "突破1星強化券30%(24星)", Probability: 29.00, MarketPrice: 0},
		{Name: "突破1星強化券50%(24星)", Probability: 19.00, MarketPrice: 0},
		{Name: "突破1星強化券100%(24星)", Probability: 14.00, MarketPrice: 0},
		{Name: "突破1星強化券30%(25星)", Probability: 20.00, MarketPrice: 0},
		{Name: "突破1星強化券50%(25星)", Probability: 9.00, MarketPrice: 0},
		{Name: "突破1星強化券100%(25星)", Probability: 4.00, MarketPrice: 0},
		{Name: "突破1星強化券30%(26星)", Probability: 3.00, MarketPrice: 0},
		{Name: "突破1星強化券50%(26星)", Probability: 2.00, MarketPrice: 0},
	},
}

// UpgradeItems 各階段的升級道具名稱
var UpgradeItems = map[int]string{
	2: "星光原石",
	3: "星光水晶",
	4: "璀璨星光",
}

// ValuableItems 需要輸入價值的道具（第一階段）
var ValuableItems = []string{
	"傳說潛在能力卷軸50%",
	"傳說潛在能力卷軸100%",
	"星力14星強化券",
	"星力15星強化券",
	"星力16星強化券",
	"星力17星強化券",
	"星力18星強化券",
	"星力19星強化券",
	"星力20星強化券",
	"星力21星強化券",
	"突破1星強化券100%(21星)",
	"突破1星強化券100%(22星)",
	"追加1星強化券30%(23星)",
	"突破1星強化券30%(23星)",
	"突破1星強化券50%(23星)",
	"突破1星強化券100%(23星)",
	"突破1星強化券30%(24星)",
	"突破1星強化券50%(24星)",
	"突破1星強化券100%(24星)",
	"突破1星強化券30%(25星)",
	"突破1星強化券50%(25星)",
	"突破1星強化券100%(25星)",
	"突破1星強化券30%(26星)",
	"突破1星強化券50%(26星)",
}

// ZeroValueItems 價值為0的道具
var ZeroValueItems = []string{
	"靈魂艾爾達碎片交換券(10個)",
	"靈魂艾爾達",
	"永遠的輪迴星火",
	"暗黑輪迴星火",
	"特別附加潛在能力賦予卷軸",
}

// LadderResult 階梯模擬結果
type LadderResult struct {
	InitialCount   int            // 初始二階錦囊數量
	Stage2Failures int            // 第2層失敗次數
	Stage3Failures int            // 第3層失敗次數
	Stage4Failures int            // 第4層失敗次數
	Stage5Success  int            // 成功到達第5階段次數
	Rewards        map[string]int // 獲得的所有獎品
}

// SimulationResult 模擬結果
type SimulationResult struct {
	DrawCount      int            // 抽獎次數
	Results        map[string]int // 各道具獲得數量
	CrystalCount   int            // 獲得玲瓏星光數量
	TheoreticalEV  float64        // 理論期望值
	ActualValue    float64        // 實際獲得價值
	TotalCost      int            // 總投入成本
}
