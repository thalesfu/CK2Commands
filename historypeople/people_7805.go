package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7805] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7805][167805] = People_167805
	HistoryPeopleMap[7805][247805] = People_247805
}
