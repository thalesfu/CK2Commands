package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6802] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6802][6802] = People_6802
	HistoryPeopleMap[6802][166802] = People_166802
	HistoryPeopleMap[6802][206802] = People_206802
}
