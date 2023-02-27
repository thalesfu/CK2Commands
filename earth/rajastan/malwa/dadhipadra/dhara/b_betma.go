package dhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博摩BetmaBarony struct {
	feud.BaseBarony
}

var BBetma博摩 feud.Barony = &博摩BetmaBarony{}

func init() {
    f := BBetma博摩.(*博摩BetmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "betma",
		TitleName: "博摩",
		TitleCode: "b_betma",
	}
}
