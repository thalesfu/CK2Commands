package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6443] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6443][6443] = People_6443
	HistoryPeopleMap[6443][166443] = People_166443
}
