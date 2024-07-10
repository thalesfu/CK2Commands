package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1812] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1812][31812] = People_31812
	HistoryPeopleMap[1812][71812] = People_71812
	HistoryPeopleMap[1812][81812] = People_81812
	HistoryPeopleMap[1812][191812] = People_191812
}
