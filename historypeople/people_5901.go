package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5901] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5901][125901] = People_125901
	HistoryPeopleMap[5901][145901] = People_145901
}
