package bambuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基法KiffaBarony struct {
	feud.BaseBarony
}

var BKiffa基法 feud.Barony = &基法KiffaBarony{}

func init() {
    f := BKiffa基法.(*基法KiffaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiffa",
		TitleName: "基法",
		TitleCode: "b_kiffa",
	}
}
