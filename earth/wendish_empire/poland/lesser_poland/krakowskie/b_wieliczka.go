package krakowskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维利奇卡WieliczkaBarony struct {
	feud.BaseBarony
}

var BWieliczka维利奇卡 feud.Barony = &维利奇卡WieliczkaBarony{}

func init() {
    f := BWieliczka维利奇卡.(*维利奇卡WieliczkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wieliczka",
		TitleName: "维利奇卡",
		TitleCode: "b_wieliczka",
	}
}
