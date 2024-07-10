package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6440] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6440][6440] = People_6440
	HistoryPeopleMap[6440][166440] = People_166440
}
