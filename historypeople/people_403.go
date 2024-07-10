package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[403] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[403][403] = People_403
	HistoryPeopleMap[403][30403] = People_30403
	HistoryPeopleMap[403][70403] = People_70403
	HistoryPeopleMap[403][160403] = People_160403
	HistoryPeopleMap[403][190403] = People_190403
	HistoryPeopleMap[403][260403] = People_260403
}
