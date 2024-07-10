package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2147] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2147][32147] = People_32147
	HistoryPeopleMap[2147][72147] = People_72147
	HistoryPeopleMap[2147][82147] = People_82147
	HistoryPeopleMap[2147][142147] = People_142147
	HistoryPeopleMap[2147][252147] = People_252147
}
