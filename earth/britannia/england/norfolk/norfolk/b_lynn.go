package norfolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林恩LynnBarony struct {
	feud.BaseBarony
}

var BLynn林恩 feud.Barony = &林恩LynnBarony{}

func init() {
	f := BLynn林恩.(*林恩LynnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lynn",
		TitleName: "林恩",
		TitleCode: "b_lynn",
	}
}
