package tenkasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波罗延柯多PalayamcottaiBarony struct {
	feud.BaseBarony
}

var BPalayamcottai波罗延柯多 feud.Barony = &波罗延柯多PalayamcottaiBarony{}

func init() {
    f := BPalayamcottai波罗延柯多.(*波罗延柯多PalayamcottaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palayamcottai",
		TitleName: "波罗延柯多",
		TitleCode: "b_palayamcottai",
	}
}
