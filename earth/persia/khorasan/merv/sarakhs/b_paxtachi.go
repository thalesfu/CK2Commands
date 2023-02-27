package sarakhs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕赫塔基PaxtachiBarony struct {
	feud.BaseBarony
}

var BPaxtachi帕赫塔基 feud.Barony = &帕赫塔基PaxtachiBarony{}

func init() {
    f := BPaxtachi帕赫塔基.(*帕赫塔基PaxtachiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paxtachi",
		TitleName: "帕赫塔基",
		TitleCode: "b_paxtachi",
	}
}
