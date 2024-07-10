package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3903] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3903][73903] = People_73903
	HistoryPeopleMap[3903][183903] = People_183903
}
