package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8443] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8443][138443] = People_138443
	HistoryPeopleMap[8443][168443] = People_168443
	HistoryPeopleMap[8443][188443] = People_188443
}
