package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5643] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5643][145643] = People_145643
	HistoryPeopleMap[5643][205643] = People_205643
}
