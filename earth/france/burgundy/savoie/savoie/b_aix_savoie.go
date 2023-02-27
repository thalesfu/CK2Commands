package savoie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾克斯Aix_savoieBarony struct {
	feud.BaseBarony
}

var BAix_savoie艾克斯 feud.Barony = &艾克斯Aix_savoieBarony{}

func init() {
    f := BAix_savoie艾克斯.(*艾克斯Aix_savoieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aix_savoie",
		TitleName: "艾克斯",
		TitleCode: "b_aix_savoie",
	}
}
