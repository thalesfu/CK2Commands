package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9999] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9999][159999] = People_159999
	HistoryPeopleMap[9999][179999] = People_179999
	HistoryPeopleMap[9999][219999] = People_219999
	HistoryPeopleMap[9999][1059999] = People_1059999
}
