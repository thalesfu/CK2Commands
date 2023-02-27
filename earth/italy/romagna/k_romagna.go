package romagna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RomagnaKingdom interface {
    feud.Kingdom
}

type 罗马涅RomagnaKingdom struct {
	feud.BaseKingdom
}

var KRomagna罗马涅 RomagnaKingdom = &罗马涅RomagnaKingdom{}

func init() {
	f := KRomagna罗马涅.(*罗马涅RomagnaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "romagna",
		TitleName: "罗马涅",
		TitleCode: "k_romagna",
		Dukes:  map[string]feud.Duke{},
	}

}
