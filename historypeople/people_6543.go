package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6543] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6543][106543] = People_106543
	HistoryPeopleMap[6543][166543] = People_166543
}
