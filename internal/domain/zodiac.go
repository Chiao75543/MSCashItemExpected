package domain

// Zodiac 生肖
type Zodiac string

const (
	Horse   Zodiac = "馬"
	Goat    Zodiac = "羊"
	Monkey  Zodiac = "猴"
	Rooster Zodiac = "雞"
	Dog     Zodiac = "狗"
	Pig     Zodiac = "豬"
	Rat     Zodiac = "鼠"
	Ox      Zodiac = "牛"
	Tiger   Zodiac = "虎"
	Rabbit  Zodiac = "兔"
	Dragon  Zodiac = "龍"
	Snake   Zodiac = "蛇"
)

// AllZodiacs 所有生肖（按機率從低到高排序）
var AllZodiacs = []Zodiac{
	Horse, Goat, Monkey, Rooster, Dog, Pig,
	Rat, Ox, Tiger, Rabbit, Dragon, Snake,
}

// ZodiacRates 生肖機率 (%)
var ZodiacRates = map[Zodiac]float64{
	Horse:   0.15,
	Goat:    0.20,
	Monkey:  0.25,
	Rooster: 0.80,
	Dog:     0.90,
	Pig:     1.00,
	Rat:     2.50,
	Ox:      3.00,
	Tiger:   3.50,
	Rabbit:  29.20,
	Dragon:  29.25,
	Snake:   29.25,
}

// BreathCollection 氣息收集（各生肖的數量）
type BreathCollection map[Zodiac]float64

// NewBreathCollection 建立空的氣息收集
func NewBreathCollection() BreathCollection {
	return make(BreathCollection)
}

// Clone 複製氣息收集
func (bc BreathCollection) Clone() BreathCollection {
	clone := make(BreathCollection)
	for k, v := range bc {
		clone[k] = v
	}
	return clone
}

// Min 取得指定生肖中的最小數量
func (bc BreathCollection) Min(zodiacs []Zodiac) float64 {
	min := bc[zodiacs[0]]
	for _, z := range zodiacs[1:] {
		if bc[z] < min {
			min = bc[z]
		}
	}
	return min
}

// Subtract 扣除指定生肖的數量
func (bc BreathCollection) Subtract(zodiacs []Zodiac, amount float64) {
	for _, z := range zodiacs {
		bc[z] -= amount
	}
}
