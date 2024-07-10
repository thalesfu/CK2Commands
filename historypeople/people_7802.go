package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7802] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7802][7802] = People_7802
	HistoryPeopleMap[7802][167802] = People_167802
	HistoryPeopleMap[7802][247802] = People_247802
}
