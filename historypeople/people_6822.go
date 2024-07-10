package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6822] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6822][6822] = People_6822
	HistoryPeopleMap[6822][166822] = People_166822
}
