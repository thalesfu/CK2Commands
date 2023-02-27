package magyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MagyarKingdom interface {
    feud.Kingdom
}

type 马扎尔MagyarKingdom struct {
	feud.BaseKingdom
}

var KMagyar马扎尔 MagyarKingdom = &马扎尔MagyarKingdom{}

func init() {
	f := KMagyar马扎尔.(*马扎尔MagyarKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "magyar",
		TitleName: "马扎尔",
		TitleCode: "k_magyar",
		Dukes:  map[string]feud.Duke{},
	}

}
