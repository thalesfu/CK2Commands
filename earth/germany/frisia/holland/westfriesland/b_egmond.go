package westfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃赫蒙德EgmondBarony struct {
	feud.BaseBarony
}

var BEgmond埃赫蒙德 feud.Barony = &埃赫蒙德EgmondBarony{}

func init() {
	f := BEgmond埃赫蒙德.(*埃赫蒙德EgmondBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "egmond",
		TitleName: "埃赫蒙德",
		TitleCode: "b_egmond",
	}
}
