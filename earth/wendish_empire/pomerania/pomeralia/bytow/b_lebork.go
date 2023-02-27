package bytow

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伦堡LeborkBarony struct {
	feud.BaseBarony
}

var BLebork伦堡 feud.Barony = &伦堡LeborkBarony{}

func init() {
    f := BLebork伦堡.(*伦堡LeborkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lebork",
		TitleName: "伦堡",
		TitleCode: "b_lebork",
	}
}
