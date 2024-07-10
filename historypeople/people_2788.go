package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2788] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2788][32788] = People_32788
	HistoryPeopleMap[2788][72788] = People_72788
	HistoryPeopleMap[2788][142788] = People_142788
}
