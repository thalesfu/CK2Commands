package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3579] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3579][33579] = People_33579
	HistoryPeopleMap[3579][73579] = People_73579
	HistoryPeopleMap[3579][83579] = People_83579
}
