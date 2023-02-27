package wielunska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维日赫拉斯WierzchlasBarony struct {
	feud.BaseBarony
}

var BWierzchlas维日赫拉斯 feud.Barony = &维日赫拉斯WierzchlasBarony{}

func init() {
    f := BWierzchlas维日赫拉斯.(*维日赫拉斯WierzchlasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wierzchlas",
		TitleName: "维日赫拉斯",
		TitleCode: "b_wierzchlas",
	}
}
