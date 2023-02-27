package skane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔比DalbyBarony struct {
	feud.BaseBarony
}

var BDalby达尔比 feud.Barony = &达尔比DalbyBarony{}

func init() {
    f := BDalby达尔比.(*达尔比DalbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dalby",
		TitleName: "达尔比",
		TitleCode: "b_dalby",
	}
}
