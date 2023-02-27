package palmyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔米拉PalmyraBarony struct {
	feud.BaseBarony
}

var BPalmyra巴尔米拉 feud.Barony = &巴尔米拉PalmyraBarony{}

func init() {
    f := BPalmyra巴尔米拉.(*巴尔米拉PalmyraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palmyra",
		TitleName: "巴尔米拉",
		TitleCode: "b_palmyra",
	}
}
