package tourraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔ToursBarony struct {
	feud.BaseBarony
}

var BTours图尔 feud.Barony = &图尔ToursBarony{}

func init() {
	f := BTours图尔.(*图尔ToursBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tours",
		TitleName: "图尔",
		TitleCode: "b_tours",
	}
}
