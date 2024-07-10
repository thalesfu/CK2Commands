package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5753] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5753][125753] = People_125753
	HistoryPeopleMap[5753][145753] = People_145753
	HistoryPeopleMap[5753][455753] = People_455753
}
