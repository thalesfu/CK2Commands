package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5801] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5801][125801] = People_125801
	HistoryPeopleMap[5801][145801] = People_145801
	HistoryPeopleMap[5801][205801] = People_205801
	HistoryPeopleMap[5801][465801] = People_465801
}
