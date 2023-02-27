package memel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克雷廷加莱KretingaleBarony struct {
	feud.BaseBarony
}

var BKretingale克雷廷加莱 feud.Barony = &克雷廷加莱KretingaleBarony{}

func init() {
    f := BKretingale克雷廷加莱.(*克雷廷加莱KretingaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kretingale",
		TitleName: "克雷廷加莱",
		TitleCode: "b_kretingale",
	}
}
