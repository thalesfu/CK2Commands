package gudbrandsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛姆LomBarony struct {
	feud.BaseBarony
}

var BLom洛姆 feud.Barony = &洛姆LomBarony{}

func init() {
    f := BLom洛姆.(*洛姆LomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lom",
		TitleName: "洛姆",
		TitleCode: "b_lom",
	}
}
