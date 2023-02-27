package artois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗斯LensBarony struct {
	feud.BaseBarony
}

var BLens朗斯 feud.Barony = &朗斯LensBarony{}

func init() {
    f := BLens朗斯.(*朗斯LensBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lens",
		TitleName: "朗斯",
		TitleCode: "b_lens",
	}
}
