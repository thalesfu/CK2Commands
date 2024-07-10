package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[220] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[220][220] = People_220
	HistoryPeopleMap[220][30220] = People_30220
	HistoryPeopleMap[220][40220] = People_40220
	HistoryPeopleMap[220][70220] = People_70220
	HistoryPeopleMap[220][170220] = People_170220
	HistoryPeopleMap[220][190220] = People_190220
	HistoryPeopleMap[220][200220] = People_200220
	HistoryPeopleMap[220][260220] = People_260220
	HistoryPeopleMap[220][480220] = People_480220
}
