package wu

import "github.com/thalesfu/CK2Commands/people"

func Hy() {
	people.Pollinate(getPeoples(), "wu")
}

func getPeoples() []people.Couple {
	var couples []people.Couple
	couples = append(couples, people.Couple{Husband: 2864881, Wife: 2851908}) // 元超 武蕊
	couples = append(couples, people.Couple{Husband: 2866561, Wife: 2836753}) // 倪元中 埃莱娜
	return couples
}
