package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4303] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4303][34303] = People_34303
	HistoryPeopleMap[4303][194303] = People_194303
}
