package serbia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/serbia/dioclea"
	"github.com/thalesfu/CK2Commands/earth/byzantium/serbia/rashka"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SerbiaKingdom interface {
    feud.Kingdom
    DDioclea迪奥克勒亚() 	dioclea.DiocleaDuke
    DRashka拉什卡() 	rashka.RashkaDuke
}

type 塞尔维亚SerbiaKingdom struct {
	feud.BaseKingdom
	迪奥克勒亚Dioclea 	dioclea.DiocleaDuke
	拉什卡Rashka 	rashka.RashkaDuke
}

func (k *塞尔维亚SerbiaKingdom) DDioclea迪奥克勒亚() dioclea.DiocleaDuke {
	return k.迪奥克勒亚Dioclea
}
    
func (k *塞尔维亚SerbiaKingdom) DRashka拉什卡() rashka.RashkaDuke {
	return k.拉什卡Rashka
}
    
var KSerbia塞尔维亚 SerbiaKingdom = &塞尔维亚SerbiaKingdom{}

func init() {
	f := KSerbia塞尔维亚.(*塞尔维亚SerbiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "serbia",
		TitleName: "塞尔维亚",
		TitleCode: "k_serbia",
		Dukes:  map[string]feud.Duke{},
	}

	f.迪奥克勒亚Dioclea = dioclea.DDioclea迪奥克勒亚
	f.迪奥克勒亚Dioclea.SetParent(f)
	
	f.拉什卡Rashka = rashka.DRashka拉什卡
	f.拉什卡Rashka.SetParent(f)
	
}
