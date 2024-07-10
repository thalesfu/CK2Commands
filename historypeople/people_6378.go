package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6378] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6378][6378] = People_6378
	HistoryPeopleMap[6378][166378] = People_166378
}
