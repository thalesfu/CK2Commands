package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1602] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1602][41602] = People_41602
	HistoryPeopleMap[1602][71602] = People_71602
	HistoryPeopleMap[1602][161602] = People_161602
	HistoryPeopleMap[1602][191602] = People_191602
}
