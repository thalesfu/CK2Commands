package persia

import (
	"github.com/thalesfu/CK2Commands/earth/persia/persia/fars"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/jibal"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/kerman"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/khozistan"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/kurdistan"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/mafaza"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PersiaKingdom interface {
    feud.Kingdom
    DFars法尔斯() 	fars.FarsDuke
    DJibal吉巴勒() 	jibal.JibalDuke
    DKerman克尔曼() 	kerman.KermanDuke
    DKhozistan胡齐斯坦() 	khozistan.KhozistanDuke
    DKurdistan库尔德斯坦() 	kurdistan.KurdistanDuke
    DMafaza玛法扎() 	mafaza.MafazaDuke
}

type 波斯PersiaKingdom struct {
	feud.BaseKingdom
	法尔斯Fars 	fars.FarsDuke
	吉巴勒Jibal 	jibal.JibalDuke
	克尔曼Kerman 	kerman.KermanDuke
	胡齐斯坦Khozistan 	khozistan.KhozistanDuke
	库尔德斯坦Kurdistan 	kurdistan.KurdistanDuke
	玛法扎Mafaza 	mafaza.MafazaDuke
}

func (k *波斯PersiaKingdom) DFars法尔斯() fars.FarsDuke {
	return k.法尔斯Fars
}
    
func (k *波斯PersiaKingdom) DJibal吉巴勒() jibal.JibalDuke {
	return k.吉巴勒Jibal
}
    
func (k *波斯PersiaKingdom) DKerman克尔曼() kerman.KermanDuke {
	return k.克尔曼Kerman
}
    
func (k *波斯PersiaKingdom) DKhozistan胡齐斯坦() khozistan.KhozistanDuke {
	return k.胡齐斯坦Khozistan
}
    
func (k *波斯PersiaKingdom) DKurdistan库尔德斯坦() kurdistan.KurdistanDuke {
	return k.库尔德斯坦Kurdistan
}
    
func (k *波斯PersiaKingdom) DMafaza玛法扎() mafaza.MafazaDuke {
	return k.玛法扎Mafaza
}
    
var KPersia波斯 PersiaKingdom = &波斯PersiaKingdom{}

func init() {
	f := KPersia波斯.(*波斯PersiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "persia",
		TitleName: "波斯",
		TitleCode: "k_persia",
		Dukes:  map[string]feud.Duke{},
	}

	f.法尔斯Fars = fars.DFars法尔斯
	f.法尔斯Fars.SetParent(f)
	
	f.吉巴勒Jibal = jibal.DJibal吉巴勒
	f.吉巴勒Jibal.SetParent(f)
	
	f.克尔曼Kerman = kerman.DKerman克尔曼
	f.克尔曼Kerman.SetParent(f)
	
	f.胡齐斯坦Khozistan = khozistan.DKhozistan胡齐斯坦
	f.胡齐斯坦Khozistan.SetParent(f)
	
	f.库尔德斯坦Kurdistan = kurdistan.DKurdistan库尔德斯坦
	f.库尔德斯坦Kurdistan.SetParent(f)
	
	f.玛法扎Mafaza = mafaza.DMafaza玛法扎
	f.玛法扎Mafaza.SetParent(f)
	
}
