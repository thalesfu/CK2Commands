package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4285] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4285][34285] = People_34285
	HistoryPeopleMap[4285][194285] = People_194285
}
