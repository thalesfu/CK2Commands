package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5303] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5303][125303] = People_125303
}
