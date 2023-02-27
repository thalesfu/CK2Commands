package westmorland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 肯德尔KendalBarony struct {
	feud.BaseBarony
}

var BKendal肯德尔 feud.Barony = &肯德尔KendalBarony{}

func init() {
    f := BKendal肯德尔.(*肯德尔KendalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kendal",
		TitleName: "肯德尔",
		TitleCode: "b_kendal",
	}
}
