package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3923] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3923][33923] = People_33923
	HistoryPeopleMap[3923][73923] = People_73923
}
