package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5998] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5998][5998] = People_5998
	HistoryPeopleMap[5998][125998] = People_125998
	HistoryPeopleMap[5998][145998] = People_145998
}
