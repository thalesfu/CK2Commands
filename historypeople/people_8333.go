package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8333] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8333][138333] = People_138333
	HistoryPeopleMap[8333][168333] = People_168333
	HistoryPeopleMap[8333][188333] = People_188333
}
