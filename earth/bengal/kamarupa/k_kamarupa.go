package kamarupa

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa/kamarupanagara"
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa/para_lauhitya"
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa/sutiya"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KamarupaKingdom interface {
    feud.Kingdom
    DKamarupanagara迦摩缕波城() 	kamarupanagara.KamarupanagaraDuke
    DPara_lauhitya波罗劳希底耶() 	para_lauhitya.Para_lauhityaDuke
    DSutiya苏底耶() 	sutiya.SutiyaDuke
}

type 迦摩缕波KamarupaKingdom struct {
	feud.BaseKingdom
	迦摩缕波城Kamarupanagara 	kamarupanagara.KamarupanagaraDuke
	波罗劳希底耶Para_lauhitya 	para_lauhitya.Para_lauhityaDuke
	苏底耶Sutiya 	sutiya.SutiyaDuke
}

func (k *迦摩缕波KamarupaKingdom) DKamarupanagara迦摩缕波城() kamarupanagara.KamarupanagaraDuke {
	return k.迦摩缕波城Kamarupanagara
}
    
func (k *迦摩缕波KamarupaKingdom) DPara_lauhitya波罗劳希底耶() para_lauhitya.Para_lauhityaDuke {
	return k.波罗劳希底耶Para_lauhitya
}
    
func (k *迦摩缕波KamarupaKingdom) DSutiya苏底耶() sutiya.SutiyaDuke {
	return k.苏底耶Sutiya
}
    
var KKamarupa迦摩缕波 KamarupaKingdom = &迦摩缕波KamarupaKingdom{}

func init() {
	f := KKamarupa迦摩缕波.(*迦摩缕波KamarupaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "kamarupa",
		TitleName: "迦摩缕波",
		TitleCode: "k_kamarupa",
		Dukes:  map[string]feud.Duke{},
	}

	f.迦摩缕波城Kamarupanagara = kamarupanagara.DKamarupanagara迦摩缕波城
	f.迦摩缕波城Kamarupanagara.SetParent(f)
	
	f.波罗劳希底耶Para_lauhitya = para_lauhitya.DPara_lauhitya波罗劳希底耶
	f.波罗劳希底耶Para_lauhitya.SetParent(f)
	
	f.苏底耶Sutiya = sutiya.DSutiya苏底耶
	f.苏底耶Sutiya.SetParent(f)
	
}
