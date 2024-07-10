package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2640] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2640][32640] = People_32640
	HistoryPeopleMap[2640][72640] = People_72640
	HistoryPeopleMap[2640][142640] = People_142640
	HistoryPeopleMap[2640][222640] = People_222640
	HistoryPeopleMap[2640][462640] = People_462640
	HistoryPeopleMap[2640][472640] = People_472640
}
