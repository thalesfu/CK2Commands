package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7950] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7950][7950] = People_7950
	HistoryPeopleMap[7950][167950] = People_167950
	HistoryPeopleMap[7950][247950] = People_247950
}
