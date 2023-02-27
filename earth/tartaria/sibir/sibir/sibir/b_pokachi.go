package sibir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波卡奇PokachiBarony struct {
	feud.BaseBarony
}

var BPokachi波卡奇 feud.Barony = &波卡奇PokachiBarony{}

func init() {
    f := BPokachi波卡奇.(*波卡奇PokachiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pokachi",
		TitleName: "波卡奇",
		TitleCode: "b_pokachi",
	}
}
