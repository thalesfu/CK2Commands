package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7474] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7474][7474] = People_7474
	HistoryPeopleMap[7474][167474] = People_167474
	HistoryPeopleMap[7474][247474] = People_247474
}
