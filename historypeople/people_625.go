package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[625] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[625][625] = People_625
	HistoryPeopleMap[625][20625] = People_20625
	HistoryPeopleMap[625][30625] = People_30625
	HistoryPeopleMap[625][180625] = People_180625
	HistoryPeopleMap[625][260625] = People_260625
}
