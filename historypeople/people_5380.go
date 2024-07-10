package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5380] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5380][485380] = People_485380
}
