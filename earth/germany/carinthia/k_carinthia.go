package carinthia

import (
	"github.com/thalesfu/CK2Commands/earth/germany/carinthia/carinthia"
	"github.com/thalesfu/CK2Commands/earth/germany/carinthia/carniola"
	"github.com/thalesfu/CK2Commands/earth/germany/carinthia/friuli"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CarinthiaKingdom interface {
	feud.Kingdom
	DCarinthia卡林西亚() carinthia.CarinthiaDuke
	DCarniola卡尔尼奥拉() carniola.CarniolaDuke
	DFriuli弗留利() friuli.FriuliDuke
}

type 卡林西亚CarinthiaKingdom struct {
	feud.BaseKingdom
	卡林西亚Carinthia carinthia.CarinthiaDuke
	卡尔尼奥拉Carniola carniola.CarniolaDuke
	弗留利Friuli     friuli.FriuliDuke
}

func (k *卡林西亚CarinthiaKingdom) DCarinthia卡林西亚() carinthia.CarinthiaDuke {
	return k.卡林西亚Carinthia
}

func (k *卡林西亚CarinthiaKingdom) DCarniola卡尔尼奥拉() carniola.CarniolaDuke {
	return k.卡尔尼奥拉Carniola
}

func (k *卡林西亚CarinthiaKingdom) DFriuli弗留利() friuli.FriuliDuke {
	return k.弗留利Friuli
}

var KCarinthia卡林西亚 CarinthiaKingdom = &卡林西亚CarinthiaKingdom{}

func init() {
	f := KCarinthia卡林西亚.(*卡林西亚CarinthiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "carinthia",
		TitleName: "卡林西亚",
		TitleCode: "k_carinthia",
		Dukes:     map[string]feud.Duke{},
	}

	f.卡林西亚Carinthia = carinthia.DCarinthia卡林西亚
	f.卡林西亚Carinthia.SetParent(f)

	f.卡尔尼奥拉Carniola = carniola.DCarniola卡尔尼奥拉
	f.卡尔尼奥拉Carniola.SetParent(f)

	f.弗留利Friuli = friuli.DFriuli弗留利
	f.弗留利Friuli.SetParent(f)

}
