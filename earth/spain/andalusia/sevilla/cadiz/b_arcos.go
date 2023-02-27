package cadiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔科斯ArcosBarony struct {
	feud.BaseBarony
}

var BArcos阿尔科斯 feud.Barony = &阿尔科斯ArcosBarony{}

func init() {
    f := BArcos阿尔科斯.(*阿尔科斯ArcosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arcos",
		TitleName: "阿尔科斯",
		TitleCode: "b_arcos",
	}
}
