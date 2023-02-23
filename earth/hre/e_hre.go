package hre

import (
	"github.com/thalesfu/CK2Commands/earth/hre/franconia"
	"github.com/thalesfu/CK2Commands/earth/hre/saxon"
	"github.com/thalesfu/CK2Commands/earth/hre/swabia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HreEmpire interface {
	feud.Empire
	KFranconia法兰克尼亚() franconia.FranconiaKingdom
	KSaxon萨克森() saxon.SaxonKingdom
	KSwabia施瓦本() swabia.SwabiaKingdom
}

type 神圣罗马帝国HreEmpire struct {
	feud.BaseEmpire
	法兰克尼亚Franconia franconia.FranconiaKingdom
	萨克森Saxon       saxon.SaxonKingdom
	施瓦本Swabia      swabia.SwabiaKingdom
}

func (e *神圣罗马帝国HreEmpire) KFranconia法兰克尼亚() franconia.FranconiaKingdom {
	return e.法兰克尼亚Franconia
}

func (e *神圣罗马帝国HreEmpire) KSaxon萨克森() saxon.SaxonKingdom {
	return e.萨克森Saxon
}

func (e *神圣罗马帝国HreEmpire) KSwabia施瓦本() swabia.SwabiaKingdom {
	return e.施瓦本Swabia
}

var EHre神圣罗马帝国 HreEmpire = &神圣罗马帝国HreEmpire{}

func init() {
	f := EHre神圣罗马帝国.(*神圣罗马帝国HreEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "hre",
		TitleName: "神圣罗马帝国",
		TitleCode: "e_hre",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.法兰克尼亚Franconia = franconia.KFranconia法兰克尼亚
	f.法兰克尼亚Franconia.SetParent(f)
	f.萨克森Saxon = saxon.KSaxon萨克森
	f.萨克森Saxon.SetParent(f)
	f.施瓦本Swabia = swabia.KSwabia施瓦本
	f.施瓦本Swabia.SetParent(f)
}
