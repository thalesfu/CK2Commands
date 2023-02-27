package moskva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯特拉IstraBarony struct {
	feud.BaseBarony
}

var BIstra伊斯特拉 feud.Barony = &伊斯特拉IstraBarony{}

func init() {
    f := BIstra伊斯特拉.(*伊斯特拉IstraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "istra",
		TitleName: "伊斯特拉",
		TitleCode: "b_istra",
	}
}
