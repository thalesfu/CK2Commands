package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3333] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3333][33333] = People_33333
	HistoryPeopleMap[3333][73333] = People_73333
	HistoryPeopleMap[3333][83333] = People_83333
	HistoryPeopleMap[3333][93333] = People_93333
}
