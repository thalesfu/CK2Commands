package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3799] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3799][73799] = People_73799
	HistoryPeopleMap[3799][93799] = People_93799
}
