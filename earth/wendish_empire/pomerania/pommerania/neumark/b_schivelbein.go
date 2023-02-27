package neumark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 席韦尔拜因SchivelbeinBarony struct {
	feud.BaseBarony
}

var BSchivelbein席韦尔拜因 feud.Barony = &席韦尔拜因SchivelbeinBarony{}

func init() {
    f := BSchivelbein席韦尔拜因.(*席韦尔拜因SchivelbeinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schivelbein",
		TitleName: "席韦尔拜因",
		TitleCode: "b_schivelbein",
	}
}
