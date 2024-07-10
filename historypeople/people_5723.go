package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5723] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5723][125723] = People_125723
	HistoryPeopleMap[5723][145723] = People_145723
}
