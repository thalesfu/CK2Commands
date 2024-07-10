package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6553] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6553][166553] = People_166553
}
