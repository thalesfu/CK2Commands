package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7979] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7979][7979] = People_7979
	HistoryPeopleMap[7979][167979] = People_167979
	HistoryPeopleMap[7979][247979] = People_247979
}
