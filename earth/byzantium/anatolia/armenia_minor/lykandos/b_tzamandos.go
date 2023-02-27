package lykandos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特泽马多斯TzamandosBarony struct {
	feud.BaseBarony
}

var BTzamandos特泽马多斯 feud.Barony = &特泽马多斯TzamandosBarony{}

func init() {
    f := BTzamandos特泽马多斯.(*特泽马多斯TzamandosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tzamandos",
		TitleName: "特泽马多斯",
		TitleCode: "b_tzamandos",
	}
}
