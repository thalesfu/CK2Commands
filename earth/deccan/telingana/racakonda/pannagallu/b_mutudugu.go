package pannagallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆图杜古MutuduguBarony struct {
	feud.BaseBarony
}

var BMutudugu穆图杜古 feud.Barony = &穆图杜古MutuduguBarony{}

func init() {
    f := BMutudugu穆图杜古.(*穆图杜古MutuduguBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mutudugu",
		TitleName: "穆图杜古",
		TitleCode: "b_mutudugu",
	}
}
