package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4201] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4201][34201] = People_34201
	HistoryPeopleMap[4201][144201] = People_144201
	HistoryPeopleMap[4201][194201] = People_194201
}
