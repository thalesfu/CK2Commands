package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2458] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2458][72458] = People_72458
	HistoryPeopleMap[2458][142458] = People_142458
}
