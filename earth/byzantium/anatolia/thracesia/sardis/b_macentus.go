package sardis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马肯托斯MacentusBarony struct {
	feud.BaseBarony
}

var BMacentus马肯托斯 feud.Barony = &马肯托斯MacentusBarony{}

func init() {
	f := BMacentus马肯托斯.(*马肯托斯MacentusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "macentus",
		TitleName: "马肯托斯",
		TitleCode: "b_macentus",
	}
}
