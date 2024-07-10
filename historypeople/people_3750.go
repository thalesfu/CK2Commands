package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3750] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3750][33750] = People_33750
	HistoryPeopleMap[3750][73750] = People_73750
	HistoryPeopleMap[3750][83750] = People_83750
	HistoryPeopleMap[3750][93750] = People_93750
	HistoryPeopleMap[3750][223750] = People_223750
}
