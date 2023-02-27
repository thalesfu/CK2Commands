package dulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 香加XiangjiaBarony struct {
	feud.BaseBarony
}

var BXiangjia香加 feud.Barony = &香加XiangjiaBarony{}

func init() {
    f := BXiangjia香加.(*香加XiangjiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xiangjia",
		TitleName: "香加",
		TitleCode: "b_xiangjia",
	}
}
