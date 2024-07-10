package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6202] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6202][6202] = People_6202
	HistoryPeopleMap[6202][146202] = People_146202
	HistoryPeopleMap[6202][166202] = People_166202
}
