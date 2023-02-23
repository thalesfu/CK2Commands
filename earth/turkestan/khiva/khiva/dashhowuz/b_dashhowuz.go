package dashhowuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达什霍武兹DashhowuzBarony struct {
	feud.BaseBarony
}

var BDashhowuz达什霍武兹 feud.Barony = &达什霍武兹DashhowuzBarony{}

func init() {
	f := BDashhowuz达什霍武兹.(*达什霍武兹DashhowuzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dashhowuz",
		TitleName: "达什霍武兹",
		TitleCode: "b_dashhowuz",
	}
}
