package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5402] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5402][125402] = People_125402
	HistoryPeopleMap[5402][455402] = People_455402
}
