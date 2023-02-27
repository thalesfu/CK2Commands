package nangqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达那Tana_bBarony struct {
	feud.BaseBarony
}

var BTana_b达那 feud.Barony = &达那Tana_bBarony{}

func init() {
    f := BTana_b达那.(*达那Tana_bBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tana_b",
		TitleName: "达那",
		TitleCode: "b_tana_b",
	}
}
