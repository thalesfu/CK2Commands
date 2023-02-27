package vendome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣阿维StavitBarony struct {
	feud.BaseBarony
}

var BStavit圣阿维 feud.Barony = &圣阿维StavitBarony{}

func init() {
    f := BStavit圣阿维.(*圣阿维StavitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stavit",
		TitleName: "圣阿维",
		TitleCode: "b_stavit",
	}
}
