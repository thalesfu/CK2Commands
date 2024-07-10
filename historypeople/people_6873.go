package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6873] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6873][6873] = People_6873
	HistoryPeopleMap[6873][166873] = People_166873
}
