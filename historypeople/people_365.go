package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[365] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[365][20365] = People_20365
	HistoryPeopleMap[365][30365] = People_30365
	HistoryPeopleMap[365][190365] = People_190365
	HistoryPeopleMap[365][200365] = People_200365
	HistoryPeopleMap[365][260365] = People_260365
}
