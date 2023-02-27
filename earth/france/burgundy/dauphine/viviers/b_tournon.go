package viviers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔农TournonBarony struct {
	feud.BaseBarony
}

var BTournon图尔农 feud.Barony = &图尔农TournonBarony{}

func init() {
    f := BTournon图尔农.(*图尔农TournonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tournon",
		TitleName: "图尔农",
		TitleCode: "b_tournon",
	}
}
