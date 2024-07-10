package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4301] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4301][34301] = People_34301
	HistoryPeopleMap[4301][194301] = People_194301
}
