package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6379] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6379][166379] = People_166379
}
