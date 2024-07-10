package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6403] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6403][6403] = People_6403
	HistoryPeopleMap[6403][166403] = People_166403
}
