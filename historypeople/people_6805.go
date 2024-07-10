package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6805] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6805][6805] = People_6805
	HistoryPeopleMap[6805][166805] = People_166805
	HistoryPeopleMap[6805][206805] = People_206805
}
