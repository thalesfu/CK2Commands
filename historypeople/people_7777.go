package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7777] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7777][7777] = People_7777
	HistoryPeopleMap[7777][167777] = People_167777
	HistoryPeopleMap[7777][247777] = People_247777
}
