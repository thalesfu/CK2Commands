package moravia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MoraviaKingdom interface {
    feud.Kingdom
}

type 大摩拉维亚MoraviaKingdom struct {
	feud.BaseKingdom
}

var KMoravia大摩拉维亚 MoraviaKingdom = &大摩拉维亚MoraviaKingdom{}

func init() {
	f := KMoravia大摩拉维亚.(*大摩拉维亚MoraviaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "moravia",
		TitleName: "大摩拉维亚",
		TitleCode: "k_moravia",
		Dukes:  map[string]feud.Duke{},
	}

}
