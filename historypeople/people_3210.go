package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3210] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3210][3210] = People_3210
	HistoryPeopleMap[3210][33210] = People_33210
	HistoryPeopleMap[3210][73210] = People_73210
	HistoryPeopleMap[3210][83210] = People_83210
	HistoryPeopleMap[3210][93210] = People_93210
	HistoryPeopleMap[3210][183210] = People_183210
	HistoryPeopleMap[3210][463210] = People_463210
}
