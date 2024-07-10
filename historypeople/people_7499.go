package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7499] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7499][7499] = People_7499
	HistoryPeopleMap[7499][167499] = People_167499
	HistoryPeopleMap[7499][247499] = People_247499
}
