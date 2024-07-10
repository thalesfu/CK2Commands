package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2200] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2200][12200] = People_12200
	HistoryPeopleMap[2200][32200] = People_32200
	HistoryPeopleMap[2200][72200] = People_72200
	HistoryPeopleMap[2200][82200] = People_82200
	HistoryPeopleMap[2200][142200] = People_142200
	HistoryPeopleMap[2200][242200] = People_242200
	HistoryPeopleMap[2200][252200] = People_252200
	HistoryPeopleMap[2200][462200] = People_462200
}
