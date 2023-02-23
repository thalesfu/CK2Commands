package papacy

import (
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/ancona"
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/ferrara"
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/latium"
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/spoleto"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PapacyKingdom interface {
	feud.Kingdom
	DAncona安科纳() ancona.AnconaDuke
	DFerrara费拉拉() ferrara.FerraraDuke
	DLatium拉丁姆() latium.LatiumDuke
	DSpoleto斯波莱托() spoleto.SpoletoDuke
}

type 教廷属邦PapacyKingdom struct {
	feud.BaseKingdom
	安科纳Ancona   ancona.AnconaDuke
	费拉拉Ferrara  ferrara.FerraraDuke
	拉丁姆Latium   latium.LatiumDuke
	斯波莱托Spoleto spoleto.SpoletoDuke
}

func (k *教廷属邦PapacyKingdom) DAncona安科纳() ancona.AnconaDuke {
	return k.安科纳Ancona
}

func (k *教廷属邦PapacyKingdom) DFerrara费拉拉() ferrara.FerraraDuke {
	return k.费拉拉Ferrara
}

func (k *教廷属邦PapacyKingdom) DLatium拉丁姆() latium.LatiumDuke {
	return k.拉丁姆Latium
}

func (k *教廷属邦PapacyKingdom) DSpoleto斯波莱托() spoleto.SpoletoDuke {
	return k.斯波莱托Spoleto
}

var KPapacy教廷属邦 PapacyKingdom = &教廷属邦PapacyKingdom{}

func init() {
	f := KPapacy教廷属邦.(*教廷属邦PapacyKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "papacy",
		TitleName: "教廷属邦",
		TitleCode: "k_papacy",
		Dukes:     map[string]feud.Duke{},
	}

	f.安科纳Ancona = ancona.DAncona安科纳
	f.安科纳Ancona.SetParent(f)

	f.费拉拉Ferrara = ferrara.DFerrara费拉拉
	f.费拉拉Ferrara.SetParent(f)

	f.拉丁姆Latium = latium.DLatium拉丁姆
	f.拉丁姆Latium.SetParent(f)

	f.斯波莱托Spoleto = spoleto.DSpoleto斯波莱托
	f.斯波莱托Spoleto.SetParent(f)

}
