package salamanca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伦布拉莱斯LumbralesBarony struct {
	feud.BaseBarony
}

var BLumbrales伦布拉莱斯 feud.Barony = &伦布拉莱斯LumbralesBarony{}

func init() {
    f := BLumbrales伦布拉莱斯.(*伦布拉莱斯LumbralesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lumbrales",
		TitleName: "伦布拉莱斯",
		TitleCode: "b_lumbrales",
	}
}
