package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4412] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4412][4412] = People_4412
	HistoryPeopleMap[4412][34412] = People_34412
	HistoryPeopleMap[4412][74412] = People_74412
	HistoryPeopleMap[4412][194412] = People_194412
}
