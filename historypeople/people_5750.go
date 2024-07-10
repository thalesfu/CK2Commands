package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5750] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5750][125750] = People_125750
	HistoryPeopleMap[5750][145750] = People_145750
	HistoryPeopleMap[5750][205750] = People_205750
	HistoryPeopleMap[5750][455750] = People_455750
}
