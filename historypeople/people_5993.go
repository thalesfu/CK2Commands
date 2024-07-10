package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5993] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5993][125993] = People_125993
	HistoryPeopleMap[5993][145993] = People_145993
}
