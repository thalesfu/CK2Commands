package murcia

import (
	"github.com/thalesfu/CK2Commands/earth/spain/valencia/murcia/almansa"
	"github.com/thalesfu/CK2Commands/earth/spain/valencia/murcia/murcia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MurciaDuke interface {
    feud.Duke
    CAlmansa阿尔曼萨() 	almansa.AlmansaCounty
    CMurcia穆尔西亚() 	murcia.MurciaCounty
}

type 穆尔西亚MurciaDuke struct {
	feud.BaseDuke
	阿尔曼萨Almansa 	almansa.AlmansaCounty
	穆尔西亚Murcia 	murcia.MurciaCounty
}

func (d *穆尔西亚MurciaDuke) CAlmansa阿尔曼萨() almansa.AlmansaCounty {
	return d.阿尔曼萨Almansa
}
    
func (d *穆尔西亚MurciaDuke) CMurcia穆尔西亚() murcia.MurciaCounty {
	return d.穆尔西亚Murcia
}
    
var DMurcia穆尔西亚 MurciaDuke = &穆尔西亚MurciaDuke{}

func init() {
	f := DMurcia穆尔西亚.(*穆尔西亚MurciaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "murcia",
		TitleName: "穆尔西亚",
		TitleCode: "d_murcia",
		Counties:  map[string]feud.County{},
	}

	f.阿尔曼萨Almansa = almansa.CAlmansa阿尔曼萨
	f.阿尔曼萨Almansa.SetParent(f)
	
	f.穆尔西亚Murcia = murcia.CMurcia穆尔西亚
	f.穆尔西亚Murcia.SetParent(f)
	
}
