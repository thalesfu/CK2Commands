package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5755] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5755][145755] = People_145755
	HistoryPeopleMap[5755][455755] = People_455755
}
