package damascus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德拉DaraaBarony struct {
	feud.BaseBarony
}

var BDaraa德拉 feud.Barony = &德拉DaraaBarony{}

func init() {
    f := BDaraa德拉.(*德拉DaraaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daraa",
		TitleName: "德拉",
		TitleCode: "b_daraa",
	}
}
