package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[990] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[990][30990] = People_30990
	HistoryPeopleMap[990][70990] = People_70990
	HistoryPeopleMap[990][260990] = People_260990
}
