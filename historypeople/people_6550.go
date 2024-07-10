package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6550] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6550][106550] = People_106550
	HistoryPeopleMap[6550][166550] = People_166550
}
