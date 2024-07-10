package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2709] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2709][32709] = People_32709
	HistoryPeopleMap[2709][72709] = People_72709
	HistoryPeopleMap[2709][142709] = People_142709
}
