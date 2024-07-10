package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5260] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5260][125260] = People_125260
	HistoryPeopleMap[5260][145260] = People_145260
	HistoryPeopleMap[5260][205260] = People_205260
	HistoryPeopleMap[5260][485260] = People_485260
}
