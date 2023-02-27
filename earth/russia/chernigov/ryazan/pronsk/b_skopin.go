package pronsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯科平SkopinBarony struct {
	feud.BaseBarony
}

var BSkopin斯科平 feud.Barony = &斯科平SkopinBarony{}

func init() {
    f := BSkopin斯科平.(*斯科平SkopinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skopin",
		TitleName: "斯科平",
		TitleCode: "b_skopin",
	}
}
