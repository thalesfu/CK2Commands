package aydhab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宰拜尔杰德ZabargadBarony struct {
	feud.BaseBarony
}

var BZabargad宰拜尔杰德 feud.Barony = &宰拜尔杰德ZabargadBarony{}

func init() {
    f := BZabargad宰拜尔杰德.(*宰拜尔杰德ZabargadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zabargad",
		TitleName: "宰拜尔杰德",
		TitleCode: "b_zabargad",
	}
}
