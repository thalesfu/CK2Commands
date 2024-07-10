package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4711] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4711][74711] = People_74711
	HistoryPeopleMap[4711][204711] = People_204711
	HistoryPeopleMap[4711][214711] = People_214711
}
