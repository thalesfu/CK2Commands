package agder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维瑟达尔VisedalBarony struct {
	feud.BaseBarony
}

var BVisedal维瑟达尔 feud.Barony = &维瑟达尔VisedalBarony{}

func init() {
    f := BVisedal维瑟达尔.(*维瑟达尔VisedalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "visedal",
		TitleName: "维瑟达尔",
		TitleCode: "b_visedal",
	}
}
