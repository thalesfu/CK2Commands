package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6061] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6061][96061] = People_96061
	HistoryPeopleMap[6061][146061] = People_146061
	HistoryPeopleMap[6061][166061] = People_166061
	HistoryPeopleMap[6061][256061] = People_256061
}
