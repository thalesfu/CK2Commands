package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1783] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1783][1783] = People_1783
	HistoryPeopleMap[1783][31783] = People_31783
	HistoryPeopleMap[1783][191783] = People_191783
}
