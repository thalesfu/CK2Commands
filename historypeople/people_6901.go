package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6901] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6901][6901] = People_6901
	HistoryPeopleMap[6901][166901] = People_166901
	HistoryPeopleMap[6901][216901] = People_216901
}
