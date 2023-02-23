package khojand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 肯德伊巴达姆KendibadamBarony struct {
	feud.BaseBarony
}

var BKendibadam肯德伊巴达姆 feud.Barony = &肯德伊巴达姆KendibadamBarony{}

func init() {
	f := BKendibadam肯德伊巴达姆.(*肯德伊巴达姆KendibadamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kendibadam",
		TitleName: "肯德伊巴达姆",
		TitleCode: "b_kendibadam",
	}
}
