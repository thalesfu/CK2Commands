package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5302] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5302][125302] = People_125302
}
