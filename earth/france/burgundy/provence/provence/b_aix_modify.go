package provence

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾克斯AixBarony struct {
	feud.BaseBarony
}

var BAix艾克斯 feud.Barony = &艾克斯AixBarony{}

func init() {
    f := BAix艾克斯.(*艾克斯AixBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aix",
		TitleName: "艾克斯",
		TitleCode: "b_aix",
	}
}
