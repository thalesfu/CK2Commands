package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6161] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6161][146161] = People_146161
	HistoryPeopleMap[6161][166161] = People_166161
}
