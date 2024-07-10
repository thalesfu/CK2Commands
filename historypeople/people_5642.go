package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5642] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5642][5642] = People_5642
	HistoryPeopleMap[5642][145642] = People_145642
}
