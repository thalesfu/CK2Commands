package ceredigion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷加伦TregaronBarony struct {
	feud.BaseBarony
}

var BTregaron特雷加伦 feud.Barony = &特雷加伦TregaronBarony{}

func init() {
    f := BTregaron特雷加伦.(*特雷加伦TregaronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tregaron",
		TitleName: "特雷加伦",
		TitleCode: "b_tregaron",
	}
}
