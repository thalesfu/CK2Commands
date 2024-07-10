package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[987] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[987][70987] = People_70987
	HistoryPeopleMap[987][260987] = People_260987
}
