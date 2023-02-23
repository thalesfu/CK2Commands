package mathura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 秣菟罗MathuraBarony struct {
	feud.BaseBarony
}

var BMathura秣菟罗 feud.Barony = &秣菟罗MathuraBarony{}

func init() {
	f := BMathura秣菟罗.(*秣菟罗MathuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mathura",
		TitleName: "秣菟罗",
		TitleCode: "b_mathura",
	}
}
