package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2777] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2777][72777] = People_72777
	HistoryPeopleMap[2777][142777] = People_142777
}
