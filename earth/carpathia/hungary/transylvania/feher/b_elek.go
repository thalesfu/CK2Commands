package feher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃莱克ElekBarony struct {
	feud.BaseBarony
}

var BElek埃莱克 feud.Barony = &埃莱克ElekBarony{}

func init() {
	f := BElek埃莱克.(*埃莱克ElekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elek",
		TitleName: "埃莱克",
		TitleCode: "b_elek",
	}
}
