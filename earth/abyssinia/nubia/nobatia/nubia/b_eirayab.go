package nubia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾拉雅布EirayabBarony struct {
	feud.BaseBarony
}

var BEirayab艾拉雅布 feud.Barony = &艾拉雅布EirayabBarony{}

func init() {
	f := BEirayab艾拉雅布.(*艾拉雅布EirayabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eirayab",
		TitleName: "艾拉雅布",
		TitleCode: "b_eirayab",
	}
}
