package burhanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃罗祇梨HalgeriBarony struct {
	feud.BaseBarony
}

var BHalgeri诃罗祇梨 feud.Barony = &诃罗祇梨HalgeriBarony{}

func init() {
    f := BHalgeri诃罗祇梨.(*诃罗祇梨HalgeriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halgeri",
		TitleName: "诃罗祇梨",
		TitleCode: "b_halgeri",
	}
}
