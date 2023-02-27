package gurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉米扬RamianBarony struct {
	feud.BaseBarony
}

var BRamian拉米扬 feud.Barony = &拉米扬RamianBarony{}

func init() {
    f := BRamian拉米扬.(*拉米扬RamianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramian",
		TitleName: "拉米扬",
		TitleCode: "b_ramian",
	}
}
