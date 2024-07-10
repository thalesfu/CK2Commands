package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4343] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4343][34343] = People_34343
	HistoryPeopleMap[4343][194343] = People_194343
}
