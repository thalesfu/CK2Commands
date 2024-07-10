package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5610] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5610][5610] = People_5610
	HistoryPeopleMap[5610][145610] = People_145610
}
