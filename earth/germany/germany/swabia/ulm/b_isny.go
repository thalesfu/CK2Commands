package ulm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯尼IsnyBarony struct {
	feud.BaseBarony
}

var BIsny伊斯尼 feud.Barony = &伊斯尼IsnyBarony{}

func init() {
    f := BIsny伊斯尼.(*伊斯尼IsnyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isny",
		TitleName: "伊斯尼",
		TitleCode: "b_isny",
	}
}
