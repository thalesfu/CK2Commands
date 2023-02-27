package bjarmia

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/bjarmia/bjarmia"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/bjarmia/north_dvina"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/bjarmia/onega_peninsula"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/bjarmia/pinega"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BjarmiaDuke interface {
    feud.Duke
    CBjarmia比亚尔米亚() 	bjarmia.BjarmiaCounty
    CNorth_dvina德维纳() 	north_dvina.North_dvinaCounty
    COnega_peninsula奥涅加半岛() 	onega_peninsula.Onega_peninsulaCounty
    CPinega库洛伊() 	pinega.PinegaCounty
}

type 比亚尔米亚BjarmiaDuke struct {
	feud.BaseDuke
	比亚尔米亚Bjarmia 	bjarmia.BjarmiaCounty
	德维纳North_dvina 	north_dvina.North_dvinaCounty
	奥涅加半岛Onega_peninsula 	onega_peninsula.Onega_peninsulaCounty
	库洛伊Pinega 	pinega.PinegaCounty
}

func (d *比亚尔米亚BjarmiaDuke) CBjarmia比亚尔米亚() bjarmia.BjarmiaCounty {
	return d.比亚尔米亚Bjarmia
}
    
func (d *比亚尔米亚BjarmiaDuke) CNorth_dvina德维纳() north_dvina.North_dvinaCounty {
	return d.德维纳North_dvina
}
    
func (d *比亚尔米亚BjarmiaDuke) COnega_peninsula奥涅加半岛() onega_peninsula.Onega_peninsulaCounty {
	return d.奥涅加半岛Onega_peninsula
}
    
func (d *比亚尔米亚BjarmiaDuke) CPinega库洛伊() pinega.PinegaCounty {
	return d.库洛伊Pinega
}
    
var DBjarmia比亚尔米亚 BjarmiaDuke = &比亚尔米亚BjarmiaDuke{}

func init() {
	f := DBjarmia比亚尔米亚.(*比亚尔米亚BjarmiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bjarmia",
		TitleName: "比亚尔米亚",
		TitleCode: "d_bjarmia",
		Counties:  map[string]feud.County{},
	}

	f.比亚尔米亚Bjarmia = bjarmia.CBjarmia比亚尔米亚
	f.比亚尔米亚Bjarmia.SetParent(f)
	
	f.德维纳North_dvina = north_dvina.CNorth_dvina德维纳
	f.德维纳North_dvina.SetParent(f)
	
	f.奥涅加半岛Onega_peninsula = onega_peninsula.COnega_peninsula奥涅加半岛
	f.奥涅加半岛Onega_peninsula.SetParent(f)
	
	f.库洛伊Pinega = pinega.CPinega库洛伊
	f.库洛伊Pinega.SetParent(f)
	
}
