package sodermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦德布鲁恩VaderbrunnBarony struct {
	feud.BaseBarony
}

var BVaderbrunn韦德布鲁恩 feud.Barony = &韦德布鲁恩VaderbrunnBarony{}

func init() {
    f := BVaderbrunn韦德布鲁恩.(*韦德布鲁恩VaderbrunnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaderbrunn",
		TitleName: "韦德布鲁恩",
		TitleCode: "b_vaderbrunn",
	}
}
