package daylam

import (
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/azerbaijan"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/dihistan"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/gilan"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/mazandaran"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/tabriz"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DaylamKingdom interface {
    feud.Kingdom
    DAzerbaijan阿塞拜疆() 	azerbaijan.AzerbaijanDuke
    DDihistan大益斯坦() 	dihistan.DihistanDuke
    DGilan吉兰() 	gilan.GilanDuke
    DMazandaran陀拔斯单() 	mazandaran.MazandaranDuke
    DTabriz大不里士() 	tabriz.TabrizDuke
}

type 德莱木DaylamKingdom struct {
	feud.BaseKingdom
	阿塞拜疆Azerbaijan 	azerbaijan.AzerbaijanDuke
	大益斯坦Dihistan 	dihistan.DihistanDuke
	吉兰Gilan 	gilan.GilanDuke
	陀拔斯单Mazandaran 	mazandaran.MazandaranDuke
	大不里士Tabriz 	tabriz.TabrizDuke
}

func (k *德莱木DaylamKingdom) DAzerbaijan阿塞拜疆() azerbaijan.AzerbaijanDuke {
	return k.阿塞拜疆Azerbaijan
}
    
func (k *德莱木DaylamKingdom) DDihistan大益斯坦() dihistan.DihistanDuke {
	return k.大益斯坦Dihistan
}
    
func (k *德莱木DaylamKingdom) DGilan吉兰() gilan.GilanDuke {
	return k.吉兰Gilan
}
    
func (k *德莱木DaylamKingdom) DMazandaran陀拔斯单() mazandaran.MazandaranDuke {
	return k.陀拔斯单Mazandaran
}
    
func (k *德莱木DaylamKingdom) DTabriz大不里士() tabriz.TabrizDuke {
	return k.大不里士Tabriz
}
    
var KDaylam德莱木 DaylamKingdom = &德莱木DaylamKingdom{}

func init() {
	f := KDaylam德莱木.(*德莱木DaylamKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "daylam",
		TitleName: "德莱木",
		TitleCode: "k_daylam",
		Dukes:  map[string]feud.Duke{},
	}

	f.阿塞拜疆Azerbaijan = azerbaijan.DAzerbaijan阿塞拜疆
	f.阿塞拜疆Azerbaijan.SetParent(f)
	
	f.大益斯坦Dihistan = dihistan.DDihistan大益斯坦
	f.大益斯坦Dihistan.SetParent(f)
	
	f.吉兰Gilan = gilan.DGilan吉兰
	f.吉兰Gilan.SetParent(f)
	
	f.陀拔斯单Mazandaran = mazandaran.DMazandaran陀拔斯单
	f.陀拔斯单Mazandaran.SetParent(f)
	
	f.大不里士Tabriz = tabriz.DTabriz大不里士
	f.大不里士Tabriz.SetParent(f)
	
}
