package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7490] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7490][7490] = People_7490
	HistoryPeopleMap[7490][167490] = People_167490
	HistoryPeopleMap[7490][247490] = People_247490
}
