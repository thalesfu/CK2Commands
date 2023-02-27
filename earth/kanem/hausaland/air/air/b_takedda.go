package air

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔凯达TakeddaBarony struct {
	feud.BaseBarony
}

var BTakedda塔凯达 feud.Barony = &塔凯达TakeddaBarony{}

func init() {
    f := BTakedda塔凯达.(*塔凯达TakeddaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "takedda",
		TitleName: "塔凯达",
		TitleCode: "b_takedda",
	}
}
