package kuban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普谢库普斯PsekupsBarony struct {
	feud.BaseBarony
}

var BPsekups普谢库普斯 feud.Barony = &普谢库普斯PsekupsBarony{}

func init() {
    f := BPsekups普谢库普斯.(*普谢库普斯PsekupsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "psekups",
		TitleName: "普谢库普斯",
		TitleCode: "b_psekups",
	}
}
