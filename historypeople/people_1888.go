package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1888] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1888][31888] = People_31888
	HistoryPeopleMap[1888][71888] = People_71888
	HistoryPeopleMap[1888][191888] = People_191888
}
