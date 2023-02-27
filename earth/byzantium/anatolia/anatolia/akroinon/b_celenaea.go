package akroinon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 刻莱奈CelenaeaBarony struct {
	feud.BaseBarony
}

var BCelenaea刻莱奈 feud.Barony = &刻莱奈CelenaeaBarony{}

func init() {
    f := BCelenaea刻莱奈.(*刻莱奈CelenaeaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "celenaea",
		TitleName: "刻莱奈",
		TitleCode: "b_celenaea",
	}
}
