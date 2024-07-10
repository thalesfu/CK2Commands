package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6699] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6699][166699] = People_166699
	HistoryPeopleMap[6699][206699] = People_206699
}
