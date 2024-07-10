package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7680] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7680][7680] = People_7680
	HistoryPeopleMap[7680][107680] = People_107680
	HistoryPeopleMap[7680][167680] = People_167680
	HistoryPeopleMap[7680][247680] = People_247680
	HistoryPeopleMap[7680][457680] = People_457680
}
