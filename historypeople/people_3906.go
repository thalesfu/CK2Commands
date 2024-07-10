package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3906] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3906][3906] = People_3906
	HistoryPeopleMap[3906][73906] = People_73906
	HistoryPeopleMap[3906][183906] = People_183906
}
