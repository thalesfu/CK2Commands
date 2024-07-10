package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5958] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5958][125958] = People_125958
	HistoryPeopleMap[5958][145958] = People_145958
}
