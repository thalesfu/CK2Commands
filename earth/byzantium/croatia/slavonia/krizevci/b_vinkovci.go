package krizevci

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温科夫齐VinkovciBarony struct {
	feud.BaseBarony
}

var BVinkovci温科夫齐 feud.Barony = &温科夫齐VinkovciBarony{}

func init() {
    f := BVinkovci温科夫齐.(*温科夫齐VinkovciBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vinkovci",
		TitleName: "温科夫齐",
		TitleCode: "b_vinkovci",
	}
}
