package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1611] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1611][1611] = People_1611
	HistoryPeopleMap[1611][71611] = People_71611
	HistoryPeopleMap[1611][161611] = People_161611
	HistoryPeopleMap[1611][191611] = People_191611
}
