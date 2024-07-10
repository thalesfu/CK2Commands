package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2432] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2432][12432] = People_12432
	HistoryPeopleMap[2432][72432] = People_72432
	HistoryPeopleMap[2432][142432] = People_142432
}
