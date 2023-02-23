package ghazna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阔悉多KhostBarony struct {
	feud.BaseBarony
}

var BKhost阔悉多 feud.Barony = &阔悉多KhostBarony{}

func init() {
	f := BKhost阔悉多.(*阔悉多KhostBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khost",
		TitleName: "阔悉多",
		TitleCode: "b_khost",
	}
}
