package thrake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希斯托斯SestosBarony struct {
	feud.BaseBarony
}

var BSestos希斯托斯 feud.Barony = &希斯托斯SestosBarony{}

func init() {
	f := BSestos希斯托斯.(*希斯托斯SestosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sestos",
		TitleName: "希斯托斯",
		TitleCode: "b_sestos",
	}
}
