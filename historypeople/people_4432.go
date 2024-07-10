package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4432] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4432][4432] = People_4432
	HistoryPeopleMap[4432][34432] = People_34432
	HistoryPeopleMap[4432][74432] = People_74432
	HistoryPeopleMap[4432][194432] = People_194432
}
