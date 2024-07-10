package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7601] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7601][7601] = People_7601
	HistoryPeopleMap[7601][107601] = People_107601
	HistoryPeopleMap[7601][167601] = People_167601
	HistoryPeopleMap[7601][247601] = People_247601
}
