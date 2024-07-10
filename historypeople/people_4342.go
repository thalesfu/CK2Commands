package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4342] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4342][34342] = People_34342
	HistoryPeopleMap[4342][194342] = People_194342
}
