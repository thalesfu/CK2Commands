package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6442] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6442][6442] = People_6442
	HistoryPeopleMap[6442][166442] = People_166442
}
