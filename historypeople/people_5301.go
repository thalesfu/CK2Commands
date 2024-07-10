package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5301] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5301][125301] = People_125301
}
