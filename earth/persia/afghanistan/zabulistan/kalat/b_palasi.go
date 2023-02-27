package kalat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕腊斯PalasiBarony struct {
	feud.BaseBarony
}

var BPalasi帕腊斯 feud.Barony = &帕腊斯PalasiBarony{}

func init() {
    f := BPalasi帕腊斯.(*帕腊斯PalasiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palasi",
		TitleName: "帕腊斯",
		TitleCode: "b_palasi",
	}
}
