package forde

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达勒DaleBarony struct {
	feud.BaseBarony
}

var BDale达勒 feud.Barony = &达勒DaleBarony{}

func init() {
    f := BDale达勒.(*达勒DaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dale",
		TitleName: "达勒",
		TitleCode: "b_dale",
	}
}
