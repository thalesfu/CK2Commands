package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7678] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7678][7678] = People_7678
	HistoryPeopleMap[7678][167678] = People_167678
	HistoryPeopleMap[7678][247678] = People_247678
}
