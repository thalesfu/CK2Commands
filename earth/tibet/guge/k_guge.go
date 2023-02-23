package guge

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/ladakh"
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/ngari"
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/purang"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GugeKingdom interface {
	feud.Kingdom
	DLadakh拉达克() ladakh.LadakhDuke
	DNgari纳里() ngari.NgariDuke
	DPurang布让() purang.PurangDuke
}

type 古格GugeKingdom struct {
	feud.BaseKingdom
	拉达克Ladakh ladakh.LadakhDuke
	纳里Ngari   ngari.NgariDuke
	布让Purang  purang.PurangDuke
}

func (k *古格GugeKingdom) DLadakh拉达克() ladakh.LadakhDuke {
	return k.拉达克Ladakh
}

func (k *古格GugeKingdom) DNgari纳里() ngari.NgariDuke {
	return k.纳里Ngari
}

func (k *古格GugeKingdom) DPurang布让() purang.PurangDuke {
	return k.布让Purang
}

var KGuge古格 GugeKingdom = &古格GugeKingdom{}

func init() {
	f := KGuge古格.(*古格GugeKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "guge",
		TitleName: "古格",
		TitleCode: "k_guge",
		Dukes:     map[string]feud.Duke{},
	}

	f.拉达克Ladakh = ladakh.DLadakh拉达克
	f.拉达克Ladakh.SetParent(f)

	f.纳里Ngari = ngari.DNgari纳里
	f.纳里Ngari.SetParent(f)

	f.布让Purang = purang.DPurang布让
	f.布让Purang.SetParent(f)

}
