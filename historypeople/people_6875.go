package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6875] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6875][6875] = People_6875
	HistoryPeopleMap[6875][166875] = People_166875
}
