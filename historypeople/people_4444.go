package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4444] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4444][4444] = People_4444
	HistoryPeopleMap[4444][34444] = People_34444
	HistoryPeopleMap[4444][74444] = People_74444
	HistoryPeopleMap[4444][194444] = People_194444
}
