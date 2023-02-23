package yuan

import (
	"CK2Commands/people"
)

func Hy() {
	people.Pollinate(getPeoples(), "yuan")
}

func getPeoples() []people.Couple {
	var couples []people.Couple
	couples = append(couples, people.Couple{Husband: 2792390, Wife: 2770581}) // 处纳 袁令姬
	couples = append(couples, people.Couple{Husband: 2787734, Wife: 2768175}) // 绹 袁英媚
	return couples
}
