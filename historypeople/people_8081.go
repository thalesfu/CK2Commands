package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8081] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8081][138081] = People_138081
	HistoryPeopleMap[8081][168081] = People_168081
	HistoryPeopleMap[8081][188081] = People_188081
	HistoryPeopleMap[8081][248081] = People_248081
}
