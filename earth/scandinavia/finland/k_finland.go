package finland

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/finland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/karelia"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/ostrobothnia"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/savonia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FinlandKingdom interface {
    feud.Kingdom
    DFinland芬兰() 	finland.FinlandDuke
    DKarelia卡累利阿() 	karelia.KareliaDuke
    DOstrobothnia东博滕() 	ostrobothnia.OstrobothniaDuke
    DSavonia萨沃尼亚() 	savonia.SavoniaDuke
}

type 芬兰FinlandKingdom struct {
	feud.BaseKingdom
	芬兰Finland 	finland.FinlandDuke
	卡累利阿Karelia 	karelia.KareliaDuke
	东博滕Ostrobothnia 	ostrobothnia.OstrobothniaDuke
	萨沃尼亚Savonia 	savonia.SavoniaDuke
}

func (k *芬兰FinlandKingdom) DFinland芬兰() finland.FinlandDuke {
	return k.芬兰Finland
}
    
func (k *芬兰FinlandKingdom) DKarelia卡累利阿() karelia.KareliaDuke {
	return k.卡累利阿Karelia
}
    
func (k *芬兰FinlandKingdom) DOstrobothnia东博滕() ostrobothnia.OstrobothniaDuke {
	return k.东博滕Ostrobothnia
}
    
func (k *芬兰FinlandKingdom) DSavonia萨沃尼亚() savonia.SavoniaDuke {
	return k.萨沃尼亚Savonia
}
    
var KFinland芬兰 FinlandKingdom = &芬兰FinlandKingdom{}

func init() {
	f := KFinland芬兰.(*芬兰FinlandKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "finland",
		TitleName: "芬兰",
		TitleCode: "k_finland",
		Dukes:  map[string]feud.Duke{},
	}

	f.芬兰Finland = finland.DFinland芬兰
	f.芬兰Finland.SetParent(f)
	
	f.卡累利阿Karelia = karelia.DKarelia卡累利阿
	f.卡累利阿Karelia.SetParent(f)
	
	f.东博滕Ostrobothnia = ostrobothnia.DOstrobothnia东博滕
	f.东博滕Ostrobothnia.SetParent(f)
	
	f.萨沃尼亚Savonia = savonia.DSavonia萨沃尼亚
	f.萨沃尼亚Savonia.SetParent(f)
	
}
