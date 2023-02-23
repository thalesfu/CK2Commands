package persia

import (
	"github.com/thalesfu/CK2Commands/earth/persia/afghanistan"
	"github.com/thalesfu/CK2Commands/earth/persia/baluchistan"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam"
	"github.com/thalesfu/CK2Commands/earth/persia/iraq"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan"
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia"
	"github.com/thalesfu/CK2Commands/earth/persia/persia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PersiaEmpire interface {
	feud.Empire
	KAfghanistan迦布罗() afghanistan.AfghanistanKingdom
	KBaluchistan昔思丹() baluchistan.BaluchistanKingdom
	KDaylam德莱木() daylam.DaylamKingdom
	KIraq伊拉克() iraq.IraqKingdom
	KKhorasan呼罗珊() khorasan.KhorasanKingdom
	KMesopotamia贾兹拉() mesopotamia.MesopotamiaKingdom
	KPersia波斯() persia.PersiaKingdom
}

type 波斯帝国PersiaEmpire struct {
	feud.BaseEmpire
	迦布罗Afghanistan afghanistan.AfghanistanKingdom
	昔思丹Baluchistan baluchistan.BaluchistanKingdom
	德莱木Daylam      daylam.DaylamKingdom
	伊拉克Iraq        iraq.IraqKingdom
	呼罗珊Khorasan    khorasan.KhorasanKingdom
	贾兹拉Mesopotamia mesopotamia.MesopotamiaKingdom
	波斯Persia       persia.PersiaKingdom
}

func (e *波斯帝国PersiaEmpire) KAfghanistan迦布罗() afghanistan.AfghanistanKingdom {
	return e.迦布罗Afghanistan
}

func (e *波斯帝国PersiaEmpire) KBaluchistan昔思丹() baluchistan.BaluchistanKingdom {
	return e.昔思丹Baluchistan
}

func (e *波斯帝国PersiaEmpire) KDaylam德莱木() daylam.DaylamKingdom {
	return e.德莱木Daylam
}

func (e *波斯帝国PersiaEmpire) KIraq伊拉克() iraq.IraqKingdom {
	return e.伊拉克Iraq
}

func (e *波斯帝国PersiaEmpire) KKhorasan呼罗珊() khorasan.KhorasanKingdom {
	return e.呼罗珊Khorasan
}

func (e *波斯帝国PersiaEmpire) KMesopotamia贾兹拉() mesopotamia.MesopotamiaKingdom {
	return e.贾兹拉Mesopotamia
}

func (e *波斯帝国PersiaEmpire) KPersia波斯() persia.PersiaKingdom {
	return e.波斯Persia
}

var EPersia波斯帝国 PersiaEmpire = &波斯帝国PersiaEmpire{}

func init() {
	f := EPersia波斯帝国.(*波斯帝国PersiaEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "persia",
		TitleName: "波斯帝国",
		TitleCode: "e_persia",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.迦布罗Afghanistan = afghanistan.KAfghanistan迦布罗
	f.迦布罗Afghanistan.SetParent(f)
	f.昔思丹Baluchistan = baluchistan.KBaluchistan昔思丹
	f.昔思丹Baluchistan.SetParent(f)
	f.德莱木Daylam = daylam.KDaylam德莱木
	f.德莱木Daylam.SetParent(f)
	f.伊拉克Iraq = iraq.KIraq伊拉克
	f.伊拉克Iraq.SetParent(f)
	f.呼罗珊Khorasan = khorasan.KKhorasan呼罗珊
	f.呼罗珊Khorasan.SetParent(f)
	f.贾兹拉Mesopotamia = mesopotamia.KMesopotamia贾兹拉
	f.贾兹拉Mesopotamia.SetParent(f)
	f.波斯Persia = persia.KPersia波斯
	f.波斯Persia.SetParent(f)
}
