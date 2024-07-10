package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4256] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4256][34256] = People_34256
	HistoryPeopleMap[4256][194256] = People_194256
}
