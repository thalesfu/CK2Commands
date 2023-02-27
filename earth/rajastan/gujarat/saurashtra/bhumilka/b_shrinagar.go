package bhumilka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室利那揭罗ShrinagarBarony struct {
	feud.BaseBarony
}

var BShrinagar室利那揭罗 feud.Barony = &室利那揭罗ShrinagarBarony{}

func init() {
    f := BShrinagar室利那揭罗.(*室利那揭罗ShrinagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shrinagar",
		TitleName: "室利那揭罗",
		TitleCode: "b_shrinagar",
	}
}
