package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4700] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4700][34700] = People_34700
	HistoryPeopleMap[4700][74700] = People_74700
	HistoryPeopleMap[4700][204700] = People_204700
	HistoryPeopleMap[4700][214700] = People_214700
}
