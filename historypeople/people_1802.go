package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1802] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1802][31802] = People_31802
	HistoryPeopleMap[1802][71802] = People_71802
	HistoryPeopleMap[1802][191802] = People_191802
	HistoryPeopleMap[1802][451802] = People_451802
	HistoryPeopleMap[1802][461802] = People_461802
}
