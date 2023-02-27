package garhwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室利那揭罗SrinagarBarony struct {
	feud.BaseBarony
}

var BSrinagar室利那揭罗 feud.Barony = &室利那揭罗SrinagarBarony{}

func init() {
    f := BSrinagar室利那揭罗.(*室利那揭罗SrinagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "srinagar",
		TitleName: "室利那揭罗",
		TitleCode: "b_srinagar",
	}
}
