package pronsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米哈伊洛夫MikhaylovBarony struct {
	feud.BaseBarony
}

var BMikhaylov米哈伊洛夫 feud.Barony = &米哈伊洛夫MikhaylovBarony{}

func init() {
	f := BMikhaylov米哈伊洛夫.(*米哈伊洛夫MikhaylovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mikhaylov",
		TitleName: "米哈伊洛夫",
		TitleCode: "b_mikhaylov",
	}
}
