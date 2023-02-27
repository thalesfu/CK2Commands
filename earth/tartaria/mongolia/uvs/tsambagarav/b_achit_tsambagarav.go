package tsambagarav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿奇特Achit_tsambagaravBarony struct {
	feud.BaseBarony
}

var BAchit_tsambagarav阿奇特 feud.Barony = &阿奇特Achit_tsambagaravBarony{}

func init() {
    f := BAchit_tsambagarav阿奇特.(*阿奇特Achit_tsambagaravBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "achit_tsambagarav",
		TitleName: "阿奇特",
		TitleCode: "b_achit_tsambagarav",
	}
}
