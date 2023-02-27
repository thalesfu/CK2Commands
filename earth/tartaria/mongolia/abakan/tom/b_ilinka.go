package tom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊利因卡IlinkaBarony struct {
	feud.BaseBarony
}

var BIlinka伊利因卡 feud.Barony = &伊利因卡IlinkaBarony{}

func init() {
    f := BIlinka伊利因卡.(*伊利因卡IlinkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ilinka",
		TitleName: "伊利因卡",
		TitleCode: "b_ilinka",
	}
}
