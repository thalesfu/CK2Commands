package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5622] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5622][145622] = People_145622
	HistoryPeopleMap[5622][205622] = People_205622
	HistoryPeopleMap[5622][455622] = People_455622
}
