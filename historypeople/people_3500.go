package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3500] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3500][3500] = People_3500
	HistoryPeopleMap[3500][33500] = People_33500
	HistoryPeopleMap[3500][73500] = People_73500
	HistoryPeopleMap[3500][83500] = People_83500
	HistoryPeopleMap[3500][93500] = People_93500
	HistoryPeopleMap[3500][183500] = People_183500
	HistoryPeopleMap[3500][223500] = People_223500
	HistoryPeopleMap[3500][453500] = People_453500
	HistoryPeopleMap[3500][473500] = People_473500
}
