package vizagipatam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘胝波梨KotipalliBarony struct {
	feud.BaseBarony
}

var BKotipalli拘胝波梨 feud.Barony = &拘胝波梨KotipalliBarony{}

func init() {
    f := BKotipalli拘胝波梨.(*拘胝波梨KotipalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotipalli",
		TitleName: "拘胝波梨",
		TitleCode: "b_kotipalli",
	}
}
