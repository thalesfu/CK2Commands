package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7883] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7883][7883] = People_7883
	HistoryPeopleMap[7883][167883] = People_167883
	HistoryPeopleMap[7883][247883] = People_247883
}
