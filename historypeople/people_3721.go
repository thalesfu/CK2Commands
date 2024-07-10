package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3721] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3721][33721] = People_33721
	HistoryPeopleMap[3721][73721] = People_73721
	HistoryPeopleMap[3721][93721] = People_93721
}
