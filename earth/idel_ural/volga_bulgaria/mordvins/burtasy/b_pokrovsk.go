package burtasy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波克罗夫斯克PokrovskBarony struct {
	feud.BaseBarony
}

var BPokrovsk波克罗夫斯克 feud.Barony = &波克罗夫斯克PokrovskBarony{}

func init() {
    f := BPokrovsk波克罗夫斯克.(*波克罗夫斯克PokrovskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pokrovsk",
		TitleName: "波克罗夫斯克",
		TitleCode: "b_pokrovsk",
	}
}
