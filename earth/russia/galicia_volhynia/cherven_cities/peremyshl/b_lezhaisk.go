package peremyshl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱扎伊斯克LezhaiskBarony struct {
	feud.BaseBarony
}

var BLezhaisk莱扎伊斯克 feud.Barony = &莱扎伊斯克LezhaiskBarony{}

func init() {
    f := BLezhaisk莱扎伊斯克.(*莱扎伊斯克LezhaiskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lezhaisk",
		TitleName: "莱扎伊斯克",
		TitleCode: "b_lezhaisk",
	}
}
