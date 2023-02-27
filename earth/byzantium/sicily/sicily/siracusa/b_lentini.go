package siracusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伦蒂尼LentiniBarony struct {
	feud.BaseBarony
}

var BLentini伦蒂尼 feud.Barony = &伦蒂尼LentiniBarony{}

func init() {
    f := BLentini伦蒂尼.(*伦蒂尼LentiniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lentini",
		TitleName: "伦蒂尼",
		TitleCode: "b_lentini",
	}
}
