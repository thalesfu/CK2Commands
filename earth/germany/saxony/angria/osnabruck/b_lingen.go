package osnabruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林根LingenBarony struct {
	feud.BaseBarony
}

var BLingen林根 feud.Barony = &林根LingenBarony{}

func init() {
	f := BLingen林根.(*林根LingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lingen",
		TitleName: "林根",
		TitleCode: "b_lingen",
	}
}
