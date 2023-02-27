package franconia_otto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Franconia_ottoKingdom interface {
    feud.Kingdom
}

type 法兰克尼亚Franconia_ottoKingdom struct {
	feud.BaseKingdom
}

var KFranconia_otto法兰克尼亚 Franconia_ottoKingdom = &法兰克尼亚Franconia_ottoKingdom{}

func init() {
	f := KFranconia_otto法兰克尼亚.(*法兰克尼亚Franconia_ottoKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "franconia_otto",
		TitleName: "法兰克尼亚",
		TitleCode: "k_franconia_otto",
		Dukes:  map[string]feud.Duke{},
	}

}
