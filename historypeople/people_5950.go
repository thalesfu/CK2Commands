package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5950] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5950][125950] = People_125950
	HistoryPeopleMap[5950][145950] = People_145950
}
