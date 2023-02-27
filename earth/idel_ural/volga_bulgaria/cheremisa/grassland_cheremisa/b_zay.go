package grassland_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎伊ZayBarony struct {
	feud.BaseBarony
}

var BZay扎伊 feud.Barony = &扎伊ZayBarony{}

func init() {
    f := BZay扎伊.(*扎伊ZayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zay",
		TitleName: "扎伊",
		TitleCode: "b_zay",
	}
}
