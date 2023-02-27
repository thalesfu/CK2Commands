package melitene

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿萨ArcaBarony struct {
	feud.BaseBarony
}

var BArca阿萨 feud.Barony = &阿萨ArcaBarony{}

func init() {
    f := BArca阿萨.(*阿萨ArcaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arca",
		TitleName: "阿萨",
		TitleCode: "b_arca",
	}
}
