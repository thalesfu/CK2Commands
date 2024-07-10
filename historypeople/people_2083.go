package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2083] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2083][12083] = People_12083
	HistoryPeopleMap[2083][72083] = People_72083
	HistoryPeopleMap[2083][82083] = People_82083
	HistoryPeopleMap[2083][142083] = People_142083
	HistoryPeopleMap[2083][252083] = People_252083
	HistoryPeopleMap[2083][412083] = People_412083
}
