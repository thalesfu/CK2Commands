package castille

import (
	"github.com/thalesfu/CK2Commands/earth/spain/castille/castilla"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CastilleKingdom interface {
    feud.Kingdom
    DCastilla卡斯蒂利亚() 	castilla.CastillaDuke
}

type 卡斯蒂利亚CastilleKingdom struct {
	feud.BaseKingdom
	卡斯蒂利亚Castilla 	castilla.CastillaDuke
}

func (k *卡斯蒂利亚CastilleKingdom) DCastilla卡斯蒂利亚() castilla.CastillaDuke {
	return k.卡斯蒂利亚Castilla
}
    
var KCastille卡斯蒂利亚 CastilleKingdom = &卡斯蒂利亚CastilleKingdom{}

func init() {
	f := KCastille卡斯蒂利亚.(*卡斯蒂利亚CastilleKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "castille",
		TitleName: "卡斯蒂利亚",
		TitleCode: "k_castille",
		Dukes:  map[string]feud.Duke{},
	}

	f.卡斯蒂利亚Castilla = castilla.DCastilla卡斯蒂利亚
	f.卡斯蒂利亚Castilla.SetParent(f)
	
}
