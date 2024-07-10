package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2048] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2048][12048] = People_12048
	HistoryPeopleMap[2048][72048] = People_72048
	HistoryPeopleMap[2048][142048] = People_142048
	HistoryPeopleMap[2048][192048] = People_192048
	HistoryPeopleMap[2048][212048] = People_212048
	HistoryPeopleMap[2048][252048] = People_252048
	HistoryPeopleMap[2048][462048] = People_462048
}
