package ghana

import (
	"github.com/thalesfu/CK2Commands/earth/mali/ghana/ghana"
	"github.com/thalesfu/CK2Commands/earth/mali/ghana/tagant"
	"github.com/thalesfu/CK2Commands/earth/mali/ghana/timbuktu"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GhanaKingdom interface {
    feud.Kingdom
    DGhana加纳() 	ghana.GhanaDuke
    DTagant塔甘特() 	tagant.TagantDuke
    DTimbuktu廷巴克图() 	timbuktu.TimbuktuDuke
}

type 加纳GhanaKingdom struct {
	feud.BaseKingdom
	加纳Ghana 	ghana.GhanaDuke
	塔甘特Tagant 	tagant.TagantDuke
	廷巴克图Timbuktu 	timbuktu.TimbuktuDuke
}

func (k *加纳GhanaKingdom) DGhana加纳() ghana.GhanaDuke {
	return k.加纳Ghana
}
    
func (k *加纳GhanaKingdom) DTagant塔甘特() tagant.TagantDuke {
	return k.塔甘特Tagant
}
    
func (k *加纳GhanaKingdom) DTimbuktu廷巴克图() timbuktu.TimbuktuDuke {
	return k.廷巴克图Timbuktu
}
    
var KGhana加纳 GhanaKingdom = &加纳GhanaKingdom{}

func init() {
	f := KGhana加纳.(*加纳GhanaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "ghana",
		TitleName: "加纳",
		TitleCode: "k_ghana",
		Dukes:  map[string]feud.Duke{},
	}

	f.加纳Ghana = ghana.DGhana加纳
	f.加纳Ghana.SetParent(f)
	
	f.塔甘特Tagant = tagant.DTagant塔甘特
	f.塔甘特Tagant.SetParent(f)
	
	f.廷巴克图Timbuktu = timbuktu.DTimbuktu廷巴克图
	f.廷巴克图Timbuktu.SetParent(f)
	
}
