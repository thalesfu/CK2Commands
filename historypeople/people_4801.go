package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4801] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4801][34801] = People_34801
	HistoryPeopleMap[4801][74801] = People_74801
	HistoryPeopleMap[4801][204801] = People_204801
	HistoryPeopleMap[4801][214801] = People_214801
}
