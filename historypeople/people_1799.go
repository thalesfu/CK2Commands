package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1799] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1799][1799] = People_1799
	HistoryPeopleMap[1799][31799] = People_31799
	HistoryPeopleMap[1799][191799] = People_191799
	HistoryPeopleMap[1799][461799] = People_461799
}
