package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1798] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1798][1798] = People_1798
	HistoryPeopleMap[1798][31798] = People_31798
	HistoryPeopleMap[1798][191798] = People_191798
	HistoryPeopleMap[1798][461798] = People_461798
}
