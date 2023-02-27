package auvergne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫扎克MozacBarony struct {
	feud.BaseBarony
}

var BMozac莫扎克 feud.Barony = &莫扎克MozacBarony{}

func init() {
    f := BMozac莫扎克.(*莫扎克MozacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mozac",
		TitleName: "莫扎克",
		TitleCode: "b_mozac",
	}
}
