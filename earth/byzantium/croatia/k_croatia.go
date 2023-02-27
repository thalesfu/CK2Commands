package croatia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/bosnia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/croatia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/dalmatia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/slavonia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/syrmia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CroatiaKingdom interface {
    feud.Kingdom
    DBosnia波斯尼亚() 	bosnia.BosniaDuke
    DCroatia克罗地亚() 	croatia.CroatiaDuke
    DDalmatia胡姆() 	dalmatia.DalmatiaDuke
    DSlavonia斯拉沃尼亚() 	slavonia.SlavoniaDuke
    DSyrmia叙尔米亚() 	syrmia.SyrmiaDuke
}

type 克罗地亚CroatiaKingdom struct {
	feud.BaseKingdom
	波斯尼亚Bosnia 	bosnia.BosniaDuke
	克罗地亚Croatia 	croatia.CroatiaDuke
	胡姆Dalmatia 	dalmatia.DalmatiaDuke
	斯拉沃尼亚Slavonia 	slavonia.SlavoniaDuke
	叙尔米亚Syrmia 	syrmia.SyrmiaDuke
}

func (k *克罗地亚CroatiaKingdom) DBosnia波斯尼亚() bosnia.BosniaDuke {
	return k.波斯尼亚Bosnia
}
    
func (k *克罗地亚CroatiaKingdom) DCroatia克罗地亚() croatia.CroatiaDuke {
	return k.克罗地亚Croatia
}
    
func (k *克罗地亚CroatiaKingdom) DDalmatia胡姆() dalmatia.DalmatiaDuke {
	return k.胡姆Dalmatia
}
    
func (k *克罗地亚CroatiaKingdom) DSlavonia斯拉沃尼亚() slavonia.SlavoniaDuke {
	return k.斯拉沃尼亚Slavonia
}
    
func (k *克罗地亚CroatiaKingdom) DSyrmia叙尔米亚() syrmia.SyrmiaDuke {
	return k.叙尔米亚Syrmia
}
    
var KCroatia克罗地亚 CroatiaKingdom = &克罗地亚CroatiaKingdom{}

func init() {
	f := KCroatia克罗地亚.(*克罗地亚CroatiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "croatia",
		TitleName: "克罗地亚",
		TitleCode: "k_croatia",
		Dukes:  map[string]feud.Duke{},
	}

	f.波斯尼亚Bosnia = bosnia.DBosnia波斯尼亚
	f.波斯尼亚Bosnia.SetParent(f)
	
	f.克罗地亚Croatia = croatia.DCroatia克罗地亚
	f.克罗地亚Croatia.SetParent(f)
	
	f.胡姆Dalmatia = dalmatia.DDalmatia胡姆
	f.胡姆Dalmatia.SetParent(f)
	
	f.斯拉沃尼亚Slavonia = slavonia.DSlavonia斯拉沃尼亚
	f.斯拉沃尼亚Slavonia.SetParent(f)
	
	f.叙尔米亚Syrmia = syrmia.DSyrmia叙尔米亚
	f.叙尔米亚Syrmia.SetParent(f)
	
}
