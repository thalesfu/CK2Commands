package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4242] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4242][34242] = People_34242
	HistoryPeopleMap[4242][144242] = People_144242
	HistoryPeopleMap[4242][194242] = People_194242
}
