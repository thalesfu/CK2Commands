package molina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔佩德雷加尔ElpedregalBarony struct {
	feud.BaseBarony
}

var BElpedregal埃尔佩德雷加尔 feud.Barony = &埃尔佩德雷加尔ElpedregalBarony{}

func init() {
	f := BElpedregal埃尔佩德雷加尔.(*埃尔佩德雷加尔ElpedregalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elpedregal",
		TitleName: "埃尔佩德雷加尔",
		TitleCode: "b_elpedregal",
	}
}
