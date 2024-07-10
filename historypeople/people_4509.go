package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4509] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4509][34509] = People_34509
	HistoryPeopleMap[4509][74509] = People_74509
	HistoryPeopleMap[4509][204509] = People_204509
}
