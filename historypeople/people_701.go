package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[701] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[701][30701] = People_30701
	HistoryPeopleMap[701][40701] = People_40701
	HistoryPeopleMap[701][70701] = People_70701
	HistoryPeopleMap[701][260701] = People_260701
	HistoryPeopleMap[701][450701] = People_450701
}
