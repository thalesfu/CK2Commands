package campulung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮特什蒂PitestiBarony struct {
	feud.BaseBarony
}

var BPitesti皮特什蒂 feud.Barony = &皮特什蒂PitestiBarony{}

func init() {
	f := BPitesti皮特什蒂.(*皮特什蒂PitestiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pitesti",
		TitleName: "皮特什蒂",
		TitleCode: "b_pitesti",
	}
}
