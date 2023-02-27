package nisibin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯提拉德维耶马德雷萨SittiradviyyemadrasaBarony struct {
	feud.BaseBarony
}

var BSittiradviyyemadrasa斯提拉德维耶马德雷萨 feud.Barony = &斯提拉德维耶马德雷萨SittiradviyyemadrasaBarony{}

func init() {
    f := BSittiradviyyemadrasa斯提拉德维耶马德雷萨.(*斯提拉德维耶马德雷萨SittiradviyyemadrasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sittiradviyyemadrasa",
		TitleName: "斯提拉德维耶马德雷萨",
		TitleCode: "b_sittiradviyyemadrasa",
	}
}
