package minden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德雷伯DrebberBarony struct {
	feud.BaseBarony
}

var BDrebber德雷伯 feud.Barony = &德雷伯DrebberBarony{}

func init() {
	f := BDrebber德雷伯.(*德雷伯DrebberBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drebber",
		TitleName: "德雷伯",
		TitleCode: "b_drebber",
	}
}
