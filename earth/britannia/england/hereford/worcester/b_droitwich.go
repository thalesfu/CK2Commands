package worcester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德罗伊特威奇DroitwichBarony struct {
	feud.BaseBarony
}

var BDroitwich德罗伊特威奇 feud.Barony = &德罗伊特威奇DroitwichBarony{}

func init() {
    f := BDroitwich德罗伊特威奇.(*德罗伊特威奇DroitwichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "droitwich",
		TitleName: "德罗伊特威奇",
		TitleCode: "b_droitwich",
	}
}
