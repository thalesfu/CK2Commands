package auvergne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多梅拉DomeyratBarony struct {
	feud.BaseBarony
}

var BDomeyrat多梅拉 feud.Barony = &多梅拉DomeyratBarony{}

func init() {
    f := BDomeyrat多梅拉.(*多梅拉DomeyratBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "domeyrat",
		TitleName: "多梅拉",
		TitleCode: "b_domeyrat",
	}
}
