package cornouaille

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗代韦内克LandevennecBarony struct {
	feud.BaseBarony
}

var BLandevennec朗代韦内克 feud.Barony = &朗代韦内克LandevennecBarony{}

func init() {
	f := BLandevennec朗代韦内克.(*朗代韦内克LandevennecBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "landevennec",
		TitleName: "朗代韦内克",
		TitleCode: "b_landevennec",
	}
}
