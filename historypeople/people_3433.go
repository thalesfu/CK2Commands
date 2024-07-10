package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3433] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3433][33433] = People_33433
	HistoryPeopleMap[3433][83433] = People_83433
	HistoryPeopleMap[3433][93433] = People_93433
}
