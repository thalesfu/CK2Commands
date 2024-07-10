package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9606] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9606][109606] = People_109606
	HistoryPeopleMap[9606][159606] = People_159606
}
