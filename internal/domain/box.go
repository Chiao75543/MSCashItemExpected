package domain

// BoxType 心願箱類型
type BoxType string

const (
	BoxSmall  BoxType = "小吉"
	BoxMedium BoxType = "中吉"
	BoxLarge  BoxType = "大吉"
	BoxSuper  BoxType = "超越"
)

// BoxRequirements 心願箱所需生肖（需各1個湊齊）
var BoxRequirements = map[BoxType][]Zodiac{
	BoxSmall:  {Rabbit, Dragon, Snake},
	BoxMedium: {Rabbit, Dragon, Snake, Tiger, Ox, Rat},
	BoxLarge:  {Rabbit, Dragon, Snake, Tiger, Ox, Rat, Pig, Dog, Rooster},
	BoxSuper:  {Rabbit, Dragon, Snake, Tiger, Ox, Rat, Pig, Dog, Rooster, Monkey, Goat, Horse},
}

// BoxPriority 心願箱優先順序（高價值優先）
var BoxPriority = []BoxType{BoxSuper, BoxLarge, BoxMedium, BoxSmall}

// BoxValues 心願箱價值
type BoxValues struct {
	Small  float64
	Medium float64
	Large  float64
	Super  float64
}

// GetValue 取得指定類型的價值
func (bv BoxValues) GetValue(boxType BoxType) float64 {
	switch boxType {
	case BoxSmall:
		return bv.Small
	case BoxMedium:
		return bv.Medium
	case BoxLarge:
		return bv.Large
	case BoxSuper:
		return bv.Super
	default:
		return 0
	}
}

// BoxCollection 心願箱收集（各類型的數量）
type BoxCollection map[BoxType]float64

// NewBoxCollection 建立空的心願箱收集
func NewBoxCollection() BoxCollection {
	return BoxCollection{
		BoxSuper:  0,
		BoxLarge:  0,
		BoxMedium: 0,
		BoxSmall:  0,
	}
}

// TotalValue 計算總價值
func (bc BoxCollection) TotalValue(values BoxValues) float64 {
	return bc[BoxSmall]*values.Small +
		bc[BoxMedium]*values.Medium +
		bc[BoxLarge]*values.Large +
		bc[BoxSuper]*values.Super
}
