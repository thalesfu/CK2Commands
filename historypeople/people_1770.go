package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1770] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1770][1770] = People_1770
	HistoryPeopleMap[1770][31770] = People_31770
	HistoryPeopleMap[1770][71770] = People_71770
	HistoryPeopleMap[1770][161770] = People_161770
	HistoryPeopleMap[1770][191770] = People_191770
}
