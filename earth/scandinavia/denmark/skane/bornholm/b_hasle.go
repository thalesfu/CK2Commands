package bornholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海斯勒HasleBarony struct {
	feud.BaseBarony
}

var BHasle海斯勒 feud.Barony = &海斯勒HasleBarony{}

func init() {
	f := BHasle海斯勒.(*海斯勒HasleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hasle",
		TitleName: "海斯勒",
		TitleCode: "b_hasle",
	}
}
