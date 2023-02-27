package tell_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔杰纳奈特TadjenanetBarony struct {
	feud.BaseBarony
}

var BTadjenanet塔杰纳奈特 feud.Barony = &塔杰纳奈特TadjenanetBarony{}

func init() {
    f := BTadjenanet塔杰纳奈特.(*塔杰纳奈特TadjenanetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tadjenanet",
		TitleName: "塔杰纳奈特",
		TitleCode: "b_tadjenanet",
	}
}
