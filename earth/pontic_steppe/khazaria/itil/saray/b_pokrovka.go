package saray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波克罗夫卡PokrovkaBarony struct {
	feud.BaseBarony
}

var BPokrovka波克罗夫卡 feud.Barony = &波克罗夫卡PokrovkaBarony{}

func init() {
    f := BPokrovka波克罗夫卡.(*波克罗夫卡PokrovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pokrovka",
		TitleName: "波克罗夫卡",
		TitleCode: "b_pokrovka",
	}
}
