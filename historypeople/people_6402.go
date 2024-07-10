package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6402] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6402][6402] = People_6402
	HistoryPeopleMap[6402][166402] = People_166402
}
