package poher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯特雷南RostrenenBarony struct {
	feud.BaseBarony
}

var BRostrenen罗斯特雷南 feud.Barony = &罗斯特雷南RostrenenBarony{}

func init() {
    f := BRostrenen罗斯特雷南.(*罗斯特雷南RostrenenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rostrenen",
		TitleName: "罗斯特雷南",
		TitleCode: "b_rostrenen",
	}
}
