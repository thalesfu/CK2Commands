package mentese

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MenteseKingdom interface {
    feud.Kingdom
}

type 门泰谢MenteseKingdom struct {
	feud.BaseKingdom
}

var KMentese门泰谢 MenteseKingdom = &门泰谢MenteseKingdom{}

func init() {
	f := KMentese门泰谢.(*门泰谢MenteseKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "mentese",
		TitleName: "门泰谢",
		TitleCode: "k_mentese",
		Dukes:  map[string]feud.Duke{},
	}

}
