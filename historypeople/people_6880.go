package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6880] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6880][6880] = People_6880
	HistoryPeopleMap[6880][166880] = People_166880
}
