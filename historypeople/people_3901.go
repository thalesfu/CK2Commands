package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3901] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3901][3901] = People_3901
	HistoryPeopleMap[3901][33901] = People_33901
	HistoryPeopleMap[3901][73901] = People_73901
	HistoryPeopleMap[3901][183901] = People_183901
}
