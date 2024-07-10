package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[499] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[499][499] = People_499
	HistoryPeopleMap[499][20499] = People_20499
	HistoryPeopleMap[499][30499] = People_30499
	HistoryPeopleMap[499][70499] = People_70499
	HistoryPeopleMap[499][190499] = People_190499
	HistoryPeopleMap[499][260499] = People_260499
}
