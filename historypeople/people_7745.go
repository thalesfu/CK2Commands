package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7745] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7745][167745] = People_167745
	HistoryPeopleMap[7745][247745] = People_247745
}
