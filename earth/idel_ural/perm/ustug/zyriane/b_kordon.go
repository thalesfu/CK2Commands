package zyriane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔东KordonBarony struct {
	feud.BaseBarony
}

var BKordon科尔东 feud.Barony = &科尔东KordonBarony{}

func init() {
    f := BKordon科尔东.(*科尔东KordonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kordon",
		TitleName: "科尔东",
		TitleCode: "b_kordon",
	}
}
