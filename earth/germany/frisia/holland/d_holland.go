package holland

import (
	"github.com/thalesfu/CK2Commands/earth/germany/frisia/holland/holland"
	"github.com/thalesfu/CK2Commands/earth/germany/frisia/holland/sticht"
	"github.com/thalesfu/CK2Commands/earth/germany/frisia/holland/westfriesland"
	"github.com/thalesfu/CK2Commands/earth/germany/frisia/holland/zeeland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HollandDuke interface {
    feud.Duke
    CHolland荷兰() 	holland.HollandCounty
    CSticht施蒂希特() 	sticht.StichtCounty
    CWestfriesland西弗里斯兰() 	westfriesland.WestfrieslandCounty
    CZeeland泽兰() 	zeeland.ZeelandCounty
}

type 荷兰HollandDuke struct {
	feud.BaseDuke
	荷兰Holland 	holland.HollandCounty
	施蒂希特Sticht 	sticht.StichtCounty
	西弗里斯兰Westfriesland 	westfriesland.WestfrieslandCounty
	泽兰Zeeland 	zeeland.ZeelandCounty
}

func (d *荷兰HollandDuke) CHolland荷兰() holland.HollandCounty {
	return d.荷兰Holland
}
    
func (d *荷兰HollandDuke) CSticht施蒂希特() sticht.StichtCounty {
	return d.施蒂希特Sticht
}
    
func (d *荷兰HollandDuke) CWestfriesland西弗里斯兰() westfriesland.WestfrieslandCounty {
	return d.西弗里斯兰Westfriesland
}
    
func (d *荷兰HollandDuke) CZeeland泽兰() zeeland.ZeelandCounty {
	return d.泽兰Zeeland
}
    
var DHolland荷兰 HollandDuke = &荷兰HollandDuke{}

func init() {
	f := DHolland荷兰.(*荷兰HollandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "holland",
		TitleName: "荷兰",
		TitleCode: "d_holland",
		Counties:  map[string]feud.County{},
	}

	f.荷兰Holland = holland.CHolland荷兰
	f.荷兰Holland.SetParent(f)
	
	f.施蒂希特Sticht = sticht.CSticht施蒂希特
	f.施蒂希特Sticht.SetParent(f)
	
	f.西弗里斯兰Westfriesland = westfriesland.CWestfriesland西弗里斯兰
	f.西弗里斯兰Westfriesland.SetParent(f)
	
	f.泽兰Zeeland = zeeland.CZeeland泽兰
	f.泽兰Zeeland.SetParent(f)
	
}
