package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3270] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3270][33270] = People_33270
	HistoryPeopleMap[3270][73270] = People_73270
	HistoryPeopleMap[3270][83270] = People_83270
	HistoryPeopleMap[3270][93270] = People_93270
}
