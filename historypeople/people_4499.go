package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4499] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4499][34499] = People_34499
	HistoryPeopleMap[4499][74499] = People_74499
	HistoryPeopleMap[4499][204499] = People_204499
}
