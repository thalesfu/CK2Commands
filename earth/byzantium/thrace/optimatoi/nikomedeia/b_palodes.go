package nikomedeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕洛德斯PalodesBarony struct {
	feud.BaseBarony
}

var BPalodes帕洛德斯 feud.Barony = &帕洛德斯PalodesBarony{}

func init() {
    f := BPalodes帕洛德斯.(*帕洛德斯PalodesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palodes",
		TitleName: "帕洛德斯",
		TitleCode: "b_palodes",
	}
}
