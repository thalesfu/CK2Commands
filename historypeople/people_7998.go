package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7998] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7998][7998] = People_7998
	HistoryPeopleMap[7998][167998] = People_167998
	HistoryPeopleMap[7998][247998] = People_247998
}
