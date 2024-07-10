package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6888] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6888][6888] = People_6888
	HistoryPeopleMap[6888][166888] = People_166888
}
