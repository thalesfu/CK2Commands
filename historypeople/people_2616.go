package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2616] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2616][32616] = People_32616
	HistoryPeopleMap[2616][72616] = People_72616
	HistoryPeopleMap[2616][142616] = People_142616
	HistoryPeopleMap[2616][462616] = People_462616
}
