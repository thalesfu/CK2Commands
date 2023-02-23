package lotharingia

import (
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/brabant"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/koln"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/trier"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LotharingiaKingdom interface {
	feud.Kingdom
	DBrabant布拉班特() brabant.BrabantDuke
	DKoln科隆() koln.KolnDuke
	DTrier特里尔() trier.TrierDuke
}

type 中法兰克LotharingiaKingdom struct {
	feud.BaseKingdom
	布拉班特Brabant brabant.BrabantDuke
	科隆Koln      koln.KolnDuke
	特里尔Trier    trier.TrierDuke
}

func (k *中法兰克LotharingiaKingdom) DBrabant布拉班特() brabant.BrabantDuke {
	return k.布拉班特Brabant
}

func (k *中法兰克LotharingiaKingdom) DKoln科隆() koln.KolnDuke {
	return k.科隆Koln
}

func (k *中法兰克LotharingiaKingdom) DTrier特里尔() trier.TrierDuke {
	return k.特里尔Trier
}

var KLotharingia中法兰克 LotharingiaKingdom = &中法兰克LotharingiaKingdom{}

func init() {
	f := KLotharingia中法兰克.(*中法兰克LotharingiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "lotharingia",
		TitleName: "中法兰克",
		TitleCode: "k_lotharingia",
		Dukes:     map[string]feud.Duke{},
	}

	f.布拉班特Brabant = brabant.DBrabant布拉班特
	f.布拉班特Brabant.SetParent(f)

	f.科隆Koln = koln.DKoln科隆
	f.科隆Koln.SetParent(f)

	f.特里尔Trier = trier.DTrier特里尔
	f.特里尔Trier.SetParent(f)

}
