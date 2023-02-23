package cadiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫雷斯JerezBarony struct {
	feud.BaseBarony
}

var BJerez赫雷斯 feud.Barony = &赫雷斯JerezBarony{}

func init() {
	f := BJerez赫雷斯.(*赫雷斯JerezBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jerez",
		TitleName: "赫雷斯",
		TitleCode: "b_jerez",
	}
}
