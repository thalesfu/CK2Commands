package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[699] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[699][70699] = People_70699
	HistoryPeopleMap[699][260699] = People_260699
	HistoryPeopleMap[699][450699] = People_450699
}
