package tsambagarav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科布多Khovd_tsambagaravBarony struct {
	feud.BaseBarony
}

var BKhovd_tsambagarav科布多 feud.Barony = &科布多Khovd_tsambagaravBarony{}

func init() {
    f := BKhovd_tsambagarav科布多.(*科布多Khovd_tsambagaravBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khovd_tsambagarav",
		TitleName: "科布多",
		TitleCode: "b_khovd_tsambagarav",
	}
}
