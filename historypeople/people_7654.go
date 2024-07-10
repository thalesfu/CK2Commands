package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7654] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7654][7654] = People_7654
	HistoryPeopleMap[7654][167654] = People_167654
	HistoryPeopleMap[7654][207654] = People_207654
	HistoryPeopleMap[7654][247654] = People_247654
}
