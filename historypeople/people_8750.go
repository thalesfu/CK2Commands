package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8750] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8750][168750] = People_168750
	HistoryPeopleMap[8750][188750] = People_188750
	HistoryPeopleMap[8750][248750] = People_248750
}
