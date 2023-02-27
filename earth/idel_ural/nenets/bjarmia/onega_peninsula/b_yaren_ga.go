package onega_peninsula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚连加Yaren_gaBarony struct {
	feud.BaseBarony
}

var BYaren_ga亚连加 feud.Barony = &亚连加Yaren_gaBarony{}

func init() {
    f := BYaren_ga亚连加.(*亚连加Yaren_gaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yaren_ga",
		TitleName: "亚连加",
		TitleCode: "b_yaren_ga",
	}
}
