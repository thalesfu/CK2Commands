package st_gallen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣加仑StgallenBarony struct {
	feud.BaseBarony
}

var BStgallen圣加仑 feud.Barony = &圣加仑StgallenBarony{}

func init() {
    f := BStgallen圣加仑.(*圣加仑StgallenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stgallen",
		TitleName: "圣加仑",
		TitleCode: "b_stgallen",
	}
}
