package hebron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德马赤DeimacharBarony struct {
	feud.BaseBarony
}

var BDeimachar德马赤 feud.Barony = &德马赤DeimacharBarony{}

func init() {
    f := BDeimachar德马赤.(*德马赤DeimacharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deimachar",
		TitleName: "德马赤",
		TitleCode: "b_deimachar",
	}
}
