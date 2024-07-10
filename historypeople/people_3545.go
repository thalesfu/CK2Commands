package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3545] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3545][33545] = People_33545
	HistoryPeopleMap[3545][73545] = People_73545
	HistoryPeopleMap[3545][83545] = People_83545
	HistoryPeopleMap[3545][93545] = People_93545
	HistoryPeopleMap[3545][203545] = People_203545
}
