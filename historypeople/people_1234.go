package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1234] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1234][41234] = People_41234
	HistoryPeopleMap[1234][71234] = People_71234
	HistoryPeopleMap[1234][91234] = People_91234
	HistoryPeopleMap[1234][191234] = People_191234
	HistoryPeopleMap[1234][251234] = People_251234
	HistoryPeopleMap[1234][261234] = People_261234
}
