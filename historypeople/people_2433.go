package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2433] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2433][12433] = People_12433
	HistoryPeopleMap[2433][72433] = People_72433
	HistoryPeopleMap[2433][142433] = People_142433
	HistoryPeopleMap[2433][212433] = People_212433
}
