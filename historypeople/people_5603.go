package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5603] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5603][125603] = People_125603
	HistoryPeopleMap[5603][145603] = People_145603
	HistoryPeopleMap[5603][205603] = People_205603
}
