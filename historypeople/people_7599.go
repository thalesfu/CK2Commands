package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7599] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7599][7599] = People_7599
	HistoryPeopleMap[7599][167599] = People_167599
	HistoryPeopleMap[7599][247599] = People_247599
}
