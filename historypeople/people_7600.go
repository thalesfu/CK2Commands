package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7600] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7600][7600] = People_7600
	HistoryPeopleMap[7600][107600] = People_107600
	HistoryPeopleMap[7600][167600] = People_167600
	HistoryPeopleMap[7600][207600] = People_207600
	HistoryPeopleMap[7600][247600] = People_247600
	HistoryPeopleMap[7600][457600] = People_457600
}
