package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1892] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1892][31892] = People_31892
	HistoryPeopleMap[1892][71892] = People_71892
	HistoryPeopleMap[1892][191892] = People_191892
}
