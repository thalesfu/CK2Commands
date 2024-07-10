package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3700] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3700][33700] = People_33700
	HistoryPeopleMap[3700][73700] = People_73700
	HistoryPeopleMap[3700][83700] = People_83700
	HistoryPeopleMap[3700][93700] = People_93700
	HistoryPeopleMap[3700][223700] = People_223700
}
