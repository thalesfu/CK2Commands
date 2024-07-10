package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4500] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4500][34500] = People_34500
	HistoryPeopleMap[4500][74500] = People_74500
	HistoryPeopleMap[4500][204500] = People_204500
	HistoryPeopleMap[4500][214500] = People_214500
	HistoryPeopleMap[4500][454500] = People_454500
}
