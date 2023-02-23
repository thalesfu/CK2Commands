package breisgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎林根ZahringenBarony struct {
	feud.BaseBarony
}

var BZahringen扎林根 feud.Barony = &扎林根ZahringenBarony{}

func init() {
	f := BZahringen扎林根.(*扎林根ZahringenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zahringen",
		TitleName: "扎林根",
		TitleCode: "b_zahringen",
	}
}
