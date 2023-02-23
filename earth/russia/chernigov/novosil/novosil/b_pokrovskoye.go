package novosil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波克罗夫斯科耶PokrovskoyeBarony struct {
	feud.BaseBarony
}

var BPokrovskoye波克罗夫斯科耶 feud.Barony = &波克罗夫斯科耶PokrovskoyeBarony{}

func init() {
	f := BPokrovskoye波克罗夫斯科耶.(*波克罗夫斯科耶PokrovskoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pokrovskoye",
		TitleName: "波克罗夫斯科耶",
		TitleCode: "b_pokrovskoye",
	}
}
