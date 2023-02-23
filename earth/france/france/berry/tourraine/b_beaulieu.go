package tourraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博略BeaulieuBarony struct {
	feud.BaseBarony
}

var BBeaulieu博略 feud.Barony = &博略BeaulieuBarony{}

func init() {
	f := BBeaulieu博略.(*博略BeaulieuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beaulieu",
		TitleName: "博略",
		TitleCode: "b_beaulieu",
	}
}
