package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4433] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4433][4433] = People_4433
	HistoryPeopleMap[4433][34433] = People_34433
	HistoryPeopleMap[4433][74433] = People_74433
	HistoryPeopleMap[4433][194433] = People_194433
}
