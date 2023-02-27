package carpathia

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CarpathiaEmpire interface {
    feud.Empire
    KDacia瓦拉吉亚() 	dacia.DaciaKingdom
    KHungary匈牙利() 	hungary.HungaryKingdom
}

type 喀尔巴阡CarpathiaEmpire struct {
	feud.BaseEmpire
	瓦拉吉亚Dacia 	dacia.DaciaKingdom
	匈牙利Hungary 	hungary.HungaryKingdom
}

func (e *喀尔巴阡CarpathiaEmpire) KDacia瓦拉吉亚() dacia.DaciaKingdom {
	return e.瓦拉吉亚Dacia
}
    
func (e *喀尔巴阡CarpathiaEmpire) KHungary匈牙利() hungary.HungaryKingdom {
	return e.匈牙利Hungary
}
    
var ECarpathia喀尔巴阡 CarpathiaEmpire = &喀尔巴阡CarpathiaEmpire{}

func init() {
	f := ECarpathia喀尔巴阡.(*喀尔巴阡CarpathiaEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "carpathia",
		TitleName: "喀尔巴阡",
		TitleCode: "e_carpathia",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.瓦拉吉亚Dacia = dacia.KDacia瓦拉吉亚
	f.瓦拉吉亚Dacia.SetParent(f)
	f.匈牙利Hungary = hungary.KHungary匈牙利
	f.匈牙利Hungary.SetParent(f)
}
