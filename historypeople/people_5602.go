package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5602] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5602][125602] = People_125602
	HistoryPeopleMap[5602][145602] = People_145602
	HistoryPeopleMap[5602][205602] = People_205602
}
