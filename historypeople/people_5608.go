package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5608] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5608][5608] = People_5608
	HistoryPeopleMap[5608][125608] = People_125608
	HistoryPeopleMap[5608][145608] = People_145608
}
