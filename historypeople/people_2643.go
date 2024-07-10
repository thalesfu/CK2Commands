package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2643] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2643][32643] = People_32643
	HistoryPeopleMap[2643][72643] = People_72643
	HistoryPeopleMap[2643][142643] = People_142643
}
