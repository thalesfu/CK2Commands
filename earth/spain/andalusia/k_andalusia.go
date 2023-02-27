package andalusia

import (
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/cordoba"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/granada"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/sevilla"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/toledo"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AndalusiaKingdom interface {
    feud.Kingdom
    DCordoba科尔多瓦() 	cordoba.CordobaDuke
    DGranada格拉纳达() 	granada.GranadaDuke
    DSevilla塞维利亚() 	sevilla.SevillaDuke
    DToledo托雷多() 	toledo.ToledoDuke
}

type 安达卢西亚AndalusiaKingdom struct {
	feud.BaseKingdom
	科尔多瓦Cordoba 	cordoba.CordobaDuke
	格拉纳达Granada 	granada.GranadaDuke
	塞维利亚Sevilla 	sevilla.SevillaDuke
	托雷多Toledo 	toledo.ToledoDuke
}

func (k *安达卢西亚AndalusiaKingdom) DCordoba科尔多瓦() cordoba.CordobaDuke {
	return k.科尔多瓦Cordoba
}
    
func (k *安达卢西亚AndalusiaKingdom) DGranada格拉纳达() granada.GranadaDuke {
	return k.格拉纳达Granada
}
    
func (k *安达卢西亚AndalusiaKingdom) DSevilla塞维利亚() sevilla.SevillaDuke {
	return k.塞维利亚Sevilla
}
    
func (k *安达卢西亚AndalusiaKingdom) DToledo托雷多() toledo.ToledoDuke {
	return k.托雷多Toledo
}
    
var KAndalusia安达卢西亚 AndalusiaKingdom = &安达卢西亚AndalusiaKingdom{}

func init() {
	f := KAndalusia安达卢西亚.(*安达卢西亚AndalusiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "andalusia",
		TitleName: "安达卢西亚",
		TitleCode: "k_andalusia",
		Dukes:  map[string]feud.Duke{},
	}

	f.科尔多瓦Cordoba = cordoba.DCordoba科尔多瓦
	f.科尔多瓦Cordoba.SetParent(f)
	
	f.格拉纳达Granada = granada.DGranada格拉纳达
	f.格拉纳达Granada.SetParent(f)
	
	f.塞维利亚Sevilla = sevilla.DSevilla塞维利亚
	f.塞维利亚Sevilla.SetParent(f)
	
	f.托雷多Toledo = toledo.DToledo托雷多
	f.托雷多Toledo.SetParent(f)
	
}
