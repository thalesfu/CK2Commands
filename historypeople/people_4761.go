package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4761] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4761][74761] = People_74761
}
