package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7083] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7083][127083] = People_127083
	HistoryPeopleMap[7083][247083] = People_247083
}
