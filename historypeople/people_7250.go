package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7250] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7250][7250] = People_7250
	HistoryPeopleMap[7250][167250] = People_167250
	HistoryPeopleMap[7250][217250] = People_217250
	HistoryPeopleMap[7250][247250] = People_247250
}
