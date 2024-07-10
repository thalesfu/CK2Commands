package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5991] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5991][125991] = People_125991
	HistoryPeopleMap[5991][145991] = People_145991
}
