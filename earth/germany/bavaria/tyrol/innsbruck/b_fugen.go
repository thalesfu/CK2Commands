package innsbruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲根FugenBarony struct {
	feud.BaseBarony
}

var BFugen菲根 feud.Barony = &菲根FugenBarony{}

func init() {
    f := BFugen菲根.(*菲根FugenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fugen",
		TitleName: "菲根",
		TitleCode: "b_fugen",
	}
}
