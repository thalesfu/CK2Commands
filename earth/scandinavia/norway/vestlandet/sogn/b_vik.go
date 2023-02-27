package sogn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维克VikBarony struct {
	feud.BaseBarony
}

var BVik维克 feud.Barony = &维克VikBarony{}

func init() {
    f := BVik维克.(*维克VikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vik",
		TitleName: "维克",
		TitleCode: "b_vik",
	}
}
