package liege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉罗什LarocheBarony struct {
	feud.BaseBarony
}

var BLaroche拉罗什 feud.Barony = &拉罗什LarocheBarony{}

func init() {
    f := BLaroche拉罗什.(*拉罗什LarocheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laroche",
		TitleName: "拉罗什",
		TitleCode: "b_laroche",
	}
}
