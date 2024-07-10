package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4300] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4300][34300] = People_34300
	HistoryPeopleMap[4300][194300] = People_194300
}
