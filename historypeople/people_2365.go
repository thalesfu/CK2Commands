package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2365] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2365][12365] = People_12365
	HistoryPeopleMap[2365][72365] = People_72365
	HistoryPeopleMap[2365][142365] = People_142365
	HistoryPeopleMap[2365][212365] = People_212365
	HistoryPeopleMap[2365][252365] = People_252365
	HistoryPeopleMap[2365][462365] = People_462365
}
