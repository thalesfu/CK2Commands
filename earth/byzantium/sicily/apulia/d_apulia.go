package apulia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/apulia/apulia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/apulia/bari"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/apulia/lecce"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ApuliaDuke interface {
    feud.Duke
    CApulia阿普利亚() 	apulia.ApuliaCounty
    CBari巴里() 	bari.BariCounty
    CLecce莱切() 	lecce.LecceCounty
}

type 阿普利亚ApuliaDuke struct {
	feud.BaseDuke
	阿普利亚Apulia 	apulia.ApuliaCounty
	巴里Bari 	bari.BariCounty
	莱切Lecce 	lecce.LecceCounty
}

func (d *阿普利亚ApuliaDuke) CApulia阿普利亚() apulia.ApuliaCounty {
	return d.阿普利亚Apulia
}
    
func (d *阿普利亚ApuliaDuke) CBari巴里() bari.BariCounty {
	return d.巴里Bari
}
    
func (d *阿普利亚ApuliaDuke) CLecce莱切() lecce.LecceCounty {
	return d.莱切Lecce
}
    
var DApulia阿普利亚 ApuliaDuke = &阿普利亚ApuliaDuke{}

func init() {
	f := DApulia阿普利亚.(*阿普利亚ApuliaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "apulia",
		TitleName: "阿普利亚",
		TitleCode: "d_apulia",
		Counties:  map[string]feud.County{},
	}

	f.阿普利亚Apulia = apulia.CApulia阿普利亚
	f.阿普利亚Apulia.SetParent(f)
	
	f.巴里Bari = bari.CBari巴里
	f.巴里Bari.SetParent(f)
	
	f.莱切Lecce = lecce.CLecce莱切
	f.莱切Lecce.SetParent(f)
	
}
