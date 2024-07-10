package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7970] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7970][7970] = People_7970
	HistoryPeopleMap[7970][167970] = People_167970
	HistoryPeopleMap[7970][247970] = People_247970
}
