package frisia

import (
	"github.com/thalesfu/CK2Commands/earth/germany/frisia/gelre"
	"github.com/thalesfu/CK2Commands/earth/germany/frisia/holland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FrisiaKingdom interface {
	feud.Kingdom
	DGelre海尔雷() gelre.GelreDuke
	DHolland荷兰() holland.HollandDuke
}

type 弗里西亚FrisiaKingdom struct {
	feud.BaseKingdom
	海尔雷Gelre  gelre.GelreDuke
	荷兰Holland holland.HollandDuke
}

func (k *弗里西亚FrisiaKingdom) DGelre海尔雷() gelre.GelreDuke {
	return k.海尔雷Gelre
}

func (k *弗里西亚FrisiaKingdom) DHolland荷兰() holland.HollandDuke {
	return k.荷兰Holland
}

var KFrisia弗里西亚 FrisiaKingdom = &弗里西亚FrisiaKingdom{}

func init() {
	f := KFrisia弗里西亚.(*弗里西亚FrisiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "frisia",
		TitleName: "弗里西亚",
		TitleCode: "k_frisia",
		Dukes:     map[string]feud.Duke{},
	}

	f.海尔雷Gelre = gelre.DGelre海尔雷
	f.海尔雷Gelre.SetParent(f)

	f.荷兰Holland = holland.DHolland荷兰
	f.荷兰Holland.SetParent(f)

}
