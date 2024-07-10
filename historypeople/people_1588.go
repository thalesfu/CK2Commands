package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1588] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1588][41588] = People_41588
	HistoryPeopleMap[1588][71588] = People_71588
	HistoryPeopleMap[1588][191588] = People_191588
}
