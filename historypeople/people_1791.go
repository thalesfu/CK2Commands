package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1791] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1791][31791] = People_31791
	HistoryPeopleMap[1791][191791] = People_191791
}
