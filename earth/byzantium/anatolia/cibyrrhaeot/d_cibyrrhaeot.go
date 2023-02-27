package cibyrrhaeot

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/cibyrrhaeot/attaleia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/cibyrrhaeot/lykia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/cibyrrhaeot/rhodos"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CibyrrhaeotDuke interface {
    feud.Duke
    CAttaleia阿塔利亚() 	attaleia.AttaleiaCounty
    CLykia吕基亚() 	lykia.LykiaCounty
    CRhodos罗德岛() 	rhodos.RhodosCounty
}

type 基比拉奥特CibyrrhaeotDuke struct {
	feud.BaseDuke
	阿塔利亚Attaleia 	attaleia.AttaleiaCounty
	吕基亚Lykia 	lykia.LykiaCounty
	罗德岛Rhodos 	rhodos.RhodosCounty
}

func (d *基比拉奥特CibyrrhaeotDuke) CAttaleia阿塔利亚() attaleia.AttaleiaCounty {
	return d.阿塔利亚Attaleia
}
    
func (d *基比拉奥特CibyrrhaeotDuke) CLykia吕基亚() lykia.LykiaCounty {
	return d.吕基亚Lykia
}
    
func (d *基比拉奥特CibyrrhaeotDuke) CRhodos罗德岛() rhodos.RhodosCounty {
	return d.罗德岛Rhodos
}
    
var DCibyrrhaeot基比拉奥特 CibyrrhaeotDuke = &基比拉奥特CibyrrhaeotDuke{}

func init() {
	f := DCibyrrhaeot基比拉奥特.(*基比拉奥特CibyrrhaeotDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "cibyrrhaeot",
		TitleName: "基比拉奥特",
		TitleCode: "d_cibyrrhaeot",
		Counties:  map[string]feud.County{},
	}

	f.阿塔利亚Attaleia = attaleia.CAttaleia阿塔利亚
	f.阿塔利亚Attaleia.SetParent(f)
	
	f.吕基亚Lykia = lykia.CLykia吕基亚
	f.吕基亚Lykia.SetParent(f)
	
	f.罗德岛Rhodos = rhodos.CRhodos罗德岛
	f.罗德岛Rhodos.SetParent(f)
	
}
