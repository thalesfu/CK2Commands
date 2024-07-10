package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2444] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2444][72444] = People_72444
	HistoryPeopleMap[2444][142444] = People_142444
}
