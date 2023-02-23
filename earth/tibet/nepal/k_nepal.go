package nepal

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/nepal/gorkha"
	"github.com/thalesfu/CK2Commands/earth/tibet/nepal/kathmandu"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NepalKingdom interface {
	feud.Kingdom
	DGorkha廓尔喀() gorkha.GorkhaDuke
	DKathmandu加德满都() kathmandu.KathmanduDuke
}

type 尼波罗NepalKingdom struct {
	feud.BaseKingdom
	廓尔喀Gorkha     gorkha.GorkhaDuke
	加德满都Kathmandu kathmandu.KathmanduDuke
}

func (k *尼波罗NepalKingdom) DGorkha廓尔喀() gorkha.GorkhaDuke {
	return k.廓尔喀Gorkha
}

func (k *尼波罗NepalKingdom) DKathmandu加德满都() kathmandu.KathmanduDuke {
	return k.加德满都Kathmandu
}

var KNepal尼波罗 NepalKingdom = &尼波罗NepalKingdom{}

func init() {
	f := KNepal尼波罗.(*尼波罗NepalKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "nepal",
		TitleName: "尼波罗",
		TitleCode: "k_nepal",
		Dukes:     map[string]feud.Duke{},
	}

	f.廓尔喀Gorkha = gorkha.DGorkha廓尔喀
	f.廓尔喀Gorkha.SetParent(f)

	f.加德满都Kathmandu = kathmandu.DKathmandu加德满都
	f.加德满都Kathmandu.SetParent(f)

}
