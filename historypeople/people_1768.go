package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1768] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1768][1768] = People_1768
	HistoryPeopleMap[1768][31768] = People_31768
	HistoryPeopleMap[1768][71768] = People_71768
	HistoryPeopleMap[1768][161768] = People_161768
	HistoryPeopleMap[1768][191768] = People_191768
}
