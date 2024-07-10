package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3745] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3745][33745] = People_33745
	HistoryPeopleMap[3745][83745] = People_83745
	HistoryPeopleMap[3745][93745] = People_93745
}
