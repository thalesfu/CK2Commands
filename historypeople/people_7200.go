package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7200] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7200][7200] = People_7200
	HistoryPeopleMap[7200][157200] = People_157200
	HistoryPeopleMap[7200][167200] = People_167200
	HistoryPeopleMap[7200][217200] = People_217200
	HistoryPeopleMap[7200][247200] = People_247200
}
