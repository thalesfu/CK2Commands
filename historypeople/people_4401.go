package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4401] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4401][4401] = People_4401
	HistoryPeopleMap[4401][34401] = People_34401
	HistoryPeopleMap[4401][74401] = People_74401
	HistoryPeopleMap[4401][194401] = People_194401
}
