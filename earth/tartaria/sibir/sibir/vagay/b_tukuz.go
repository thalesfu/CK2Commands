package vagay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图库兹TukuzBarony struct {
	feud.BaseBarony
}

var BTukuz图库兹 feud.Barony = &图库兹TukuzBarony{}

func init() {
    f := BTukuz图库兹.(*图库兹TukuzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tukuz",
		TitleName: "图库兹",
		TitleCode: "b_tukuz",
	}
}
