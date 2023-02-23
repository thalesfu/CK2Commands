package fu

import (
	"CK2Commands/people"
)

func Hy() {
	people.Pollinate(getPeoples(), "fu")
}

func getPeoples() []people.Couple {
	var couples []people.Couple
	couples = append(couples, people.Couple{Husband: 2786678, Wife: 2789921}) // 傅钦望 农娜.特里哈雷斯
	couples = append(couples, people.Couple{Husband: 2792626, Wife: 2791021}) // 傅威 阿德尔海德.罗丁
	couples = append(couples, people.Couple{Husband: 2788850, Wife: 2783966}) // 傅守礼 傅少姬
	return couples
}
