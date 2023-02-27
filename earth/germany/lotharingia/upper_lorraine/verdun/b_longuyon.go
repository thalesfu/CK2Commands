package verdun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 隆吉永LonguyonBarony struct {
	feud.BaseBarony
}

var BLonguyon隆吉永 feud.Barony = &隆吉永LonguyonBarony{}

func init() {
    f := BLonguyon隆吉永.(*隆吉永LonguyonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "longuyon",
		TitleName: "隆吉永",
		TitleCode: "b_longuyon",
	}
}
