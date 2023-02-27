package senlis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁瓦西RoissyBarony struct {
	feud.BaseBarony
}

var BRoissy鲁瓦西 feud.Barony = &鲁瓦西RoissyBarony{}

func init() {
    f := BRoissy鲁瓦西.(*鲁瓦西RoissyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roissy",
		TitleName: "鲁瓦西",
		TitleCode: "b_roissy",
	}
}
