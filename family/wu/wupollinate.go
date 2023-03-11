package wu

import "github.com/thalesfu/CK2Commands/people"

func Pollinate() {
	people.Pollinate(getPollinatePeople(), "wu")
}

func getPollinatePeople() []people.Couple {
	var couples []people.Couple
	couples = append(couples, people.Couple{Husband: 2864881, Wife: 2851908})  // 元超 武蕊
	couples = append(couples, people.Couple{Husband: 2866561, Wife: 2836753})  // 倪元中 埃莱娜
	couples = append(couples, people.Couple{Husband: M方庆_守礼II, Wife: 2833867}) // 武方庆 殷女节
	couples = append(couples, people.Couple{Husband: 2830314, Wife: 2843890})  // 武见素 傅少姬
	couples = append(couples, people.Couple{Husband: 2853944, Wife: 2847665})  // 武三思 尼伯龙根 乌特鲁格萨
	couples = append(couples, people.Couple{Husband: 2846817, Wife: 2846879})  // 武麟 兰迪娜
	couples = append(couples, people.Couple{Husband: 2849641, Wife: 2855851})  // 武覃 阿图瓦
	couples = append(couples, people.Couple{Husband: 2851226, Wife: 2874617})  // 武德真 燕阿娥
	couples = append(couples, people.Couple{Husband: M炎, Wife: 2869818})       // 武德真 燕阿娥
	couples = append(couples, people.Couple{Husband: 2874990, Wife: 2854503})  // 云冕 武贵兰
	couples = append(couples, people.Couple{Husband: 2874999, Wife: 2859274})  // 黄埔安石 武瑚儿
	return couples
}
