package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1803] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1803][31803] = People_31803
	HistoryPeopleMap[1803][71803] = People_71803
	HistoryPeopleMap[1803][191803] = People_191803
	HistoryPeopleMap[1803][451803] = People_451803
	HistoryPeopleMap[1803][461803] = People_461803
}
