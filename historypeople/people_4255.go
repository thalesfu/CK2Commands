package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4255] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4255][34255] = People_34255
	HistoryPeopleMap[4255][194255] = People_194255
}
