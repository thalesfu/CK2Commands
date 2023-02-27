package priluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅任NizhynBarony struct {
	feud.BaseBarony
}

var BNizhyn涅任 feud.Barony = &涅任NizhynBarony{}

func init() {
    f := BNizhyn涅任.(*涅任NizhynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nizhyn",
		TitleName: "涅任",
		TitleCode: "b_nizhyn",
	}
}
