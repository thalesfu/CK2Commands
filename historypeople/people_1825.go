package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1825] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1825][31825] = People_31825
	HistoryPeopleMap[1825][71825] = People_71825
	HistoryPeopleMap[1825][81825] = People_81825
	HistoryPeopleMap[1825][191825] = People_191825
}
