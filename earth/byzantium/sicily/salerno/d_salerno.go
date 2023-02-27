package salerno

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/salerno/amalfi"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/salerno/salerno"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/salerno/taranto"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SalernoDuke interface {
    feud.Duke
    CAmalfi阿马尔菲() 	amalfi.AmalfiCounty
    CSalerno萨莱诺() 	salerno.SalernoCounty
    CTaranto塔兰托() 	taranto.TarantoCounty
}

type 萨莱诺SalernoDuke struct {
	feud.BaseDuke
	阿马尔菲Amalfi 	amalfi.AmalfiCounty
	萨莱诺Salerno 	salerno.SalernoCounty
	塔兰托Taranto 	taranto.TarantoCounty
}

func (d *萨莱诺SalernoDuke) CAmalfi阿马尔菲() amalfi.AmalfiCounty {
	return d.阿马尔菲Amalfi
}
    
func (d *萨莱诺SalernoDuke) CSalerno萨莱诺() salerno.SalernoCounty {
	return d.萨莱诺Salerno
}
    
func (d *萨莱诺SalernoDuke) CTaranto塔兰托() taranto.TarantoCounty {
	return d.塔兰托Taranto
}
    
var DSalerno萨莱诺 SalernoDuke = &萨莱诺SalernoDuke{}

func init() {
	f := DSalerno萨莱诺.(*萨莱诺SalernoDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "salerno",
		TitleName: "萨莱诺",
		TitleCode: "d_salerno",
		Counties:  map[string]feud.County{},
	}

	f.阿马尔菲Amalfi = amalfi.CAmalfi阿马尔菲
	f.阿马尔菲Amalfi.SetParent(f)
	
	f.萨莱诺Salerno = salerno.CSalerno萨莱诺
	f.萨莱诺Salerno.SetParent(f)
	
	f.塔兰托Taranto = taranto.CTaranto塔兰托
	f.塔兰托Taranto.SetParent(f)
	
}
