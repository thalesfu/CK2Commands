package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1591] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1591][41591] = People_41591
	HistoryPeopleMap[1591][71591] = People_71591
	HistoryPeopleMap[1591][191591] = People_191591
}
