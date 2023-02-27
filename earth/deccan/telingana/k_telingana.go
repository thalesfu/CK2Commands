package telingana

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/telingana/racakonda"
	"github.com/thalesfu/CK2Commands/earth/deccan/telingana/warangal"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TelinganaKingdom interface {
    feud.Kingdom
    DRacakonda罗遮军荼() 	racakonda.RacakondaDuke
    DWarangal婆浪伽罗() 	warangal.WarangalDuke
}

type 得楞伽那TelinganaKingdom struct {
	feud.BaseKingdom
	罗遮军荼Racakonda 	racakonda.RacakondaDuke
	婆浪伽罗Warangal 	warangal.WarangalDuke
}

func (k *得楞伽那TelinganaKingdom) DRacakonda罗遮军荼() racakonda.RacakondaDuke {
	return k.罗遮军荼Racakonda
}
    
func (k *得楞伽那TelinganaKingdom) DWarangal婆浪伽罗() warangal.WarangalDuke {
	return k.婆浪伽罗Warangal
}
    
var KTelingana得楞伽那 TelinganaKingdom = &得楞伽那TelinganaKingdom{}

func init() {
	f := KTelingana得楞伽那.(*得楞伽那TelinganaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "telingana",
		TitleName: "得楞伽那",
		TitleCode: "k_telingana",
		Dukes:  map[string]feud.Duke{},
	}

	f.罗遮军荼Racakonda = racakonda.DRacakonda罗遮军荼
	f.罗遮军荼Racakonda.SetParent(f)
	
	f.婆浪伽罗Warangal = warangal.DWarangal婆浪伽罗
	f.婆浪伽罗Warangal.SetParent(f)
	
}
