package chur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达沃斯DavosBarony struct {
	feud.BaseBarony
}

var BDavos达沃斯 feud.Barony = &达沃斯DavosBarony{}

func init() {
    f := BDavos达沃斯.(*达沃斯DavosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "davos",
		TitleName: "达沃斯",
		TitleCode: "b_davos",
	}
}
