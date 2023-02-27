package uchturpan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 高车田GaochetianBarony struct {
	feud.BaseBarony
}

var BGaochetian高车田 feud.Barony = &高车田GaochetianBarony{}

func init() {
    f := BGaochetian高车田.(*高车田GaochetianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gaochetian",
		TitleName: "高车田",
		TitleCode: "b_gaochetian",
	}
}
