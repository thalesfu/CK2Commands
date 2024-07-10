package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7999] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7999][7999] = People_7999
	HistoryPeopleMap[7999][167999] = People_167999
	HistoryPeopleMap[7999][247999] = People_247999
}
