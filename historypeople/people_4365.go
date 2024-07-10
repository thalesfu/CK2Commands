package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4365] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4365][4365] = People_4365
	HistoryPeopleMap[4365][34365] = People_34365
	HistoryPeopleMap[4365][194365] = People_194365
}
