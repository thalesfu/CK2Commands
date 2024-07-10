package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2451] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2451][72451] = People_72451
	HistoryPeopleMap[2451][142451] = People_142451
}
