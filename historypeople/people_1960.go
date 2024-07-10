package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1960] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1960][71960] = People_71960
	HistoryPeopleMap[1960][161960] = People_161960
	HistoryPeopleMap[1960][191960] = People_191960
	HistoryPeopleMap[1960][451960] = People_451960
}
