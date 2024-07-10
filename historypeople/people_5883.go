package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5883] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5883][145883] = People_145883
	HistoryPeopleMap[5883][205883] = People_205883
}
