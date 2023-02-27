package severnaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波克恰PokchaBarony struct {
	feud.BaseBarony
}

var BPokcha波克恰 feud.Barony = &波克恰PokchaBarony{}

func init() {
    f := BPokcha波克恰.(*波克恰PokchaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pokcha",
		TitleName: "波克恰",
		TitleCode: "b_pokcha",
	}
}
