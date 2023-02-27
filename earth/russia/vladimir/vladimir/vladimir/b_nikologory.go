package vladimir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼科洛戈雷NikologoryBarony struct {
	feud.BaseBarony
}

var BNikologory尼科洛戈雷 feud.Barony = &尼科洛戈雷NikologoryBarony{}

func init() {
    f := BNikologory尼科洛戈雷.(*尼科洛戈雷NikologoryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nikologory",
		TitleName: "尼科洛戈雷",
		TitleCode: "b_nikologory",
	}
}
