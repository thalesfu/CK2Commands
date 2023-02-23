package satyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室陵伽伐罗补拘吒SrungavarapukotaBarony struct {
	feud.BaseBarony
}

var BSrungavarapukota室陵伽伐罗补拘吒 feud.Barony = &室陵伽伐罗补拘吒SrungavarapukotaBarony{}

func init() {
	f := BSrungavarapukota室陵伽伐罗补拘吒.(*室陵伽伐罗补拘吒SrungavarapukotaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "srungavarapukota",
		TitleName: "室陵伽伐罗补拘吒",
		TitleCode: "b_srungavarapukota",
	}
}
