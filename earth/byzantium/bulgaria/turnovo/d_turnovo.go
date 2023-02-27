package turnovo

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/turnovo/dorostotum"
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/turnovo/nikopolis"
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/turnovo/serdica"
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/turnovo/tyrnovo"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TurnovoDuke interface {
    feud.Duke
    CDorostotum杜罗斯托鲁姆() 	dorostotum.DorostotumCounty
    CNikopolis尼科波利斯() 	nikopolis.NikopolisCounty
    CSerdica塞尔迪卡() 	serdica.SerdicaCounty
    CTyrnovo特尔诺沃() 	tyrnovo.TyrnovoCounty
}

type 默西亚TurnovoDuke struct {
	feud.BaseDuke
	杜罗斯托鲁姆Dorostotum 	dorostotum.DorostotumCounty
	尼科波利斯Nikopolis 	nikopolis.NikopolisCounty
	塞尔迪卡Serdica 	serdica.SerdicaCounty
	特尔诺沃Tyrnovo 	tyrnovo.TyrnovoCounty
}

func (d *默西亚TurnovoDuke) CDorostotum杜罗斯托鲁姆() dorostotum.DorostotumCounty {
	return d.杜罗斯托鲁姆Dorostotum
}
    
func (d *默西亚TurnovoDuke) CNikopolis尼科波利斯() nikopolis.NikopolisCounty {
	return d.尼科波利斯Nikopolis
}
    
func (d *默西亚TurnovoDuke) CSerdica塞尔迪卡() serdica.SerdicaCounty {
	return d.塞尔迪卡Serdica
}
    
func (d *默西亚TurnovoDuke) CTyrnovo特尔诺沃() tyrnovo.TyrnovoCounty {
	return d.特尔诺沃Tyrnovo
}
    
var DTurnovo默西亚 TurnovoDuke = &默西亚TurnovoDuke{}

func init() {
	f := DTurnovo默西亚.(*默西亚TurnovoDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "turnovo",
		TitleName: "默西亚",
		TitleCode: "d_turnovo",
		Counties:  map[string]feud.County{},
	}

	f.杜罗斯托鲁姆Dorostotum = dorostotum.CDorostotum杜罗斯托鲁姆
	f.杜罗斯托鲁姆Dorostotum.SetParent(f)
	
	f.尼科波利斯Nikopolis = nikopolis.CNikopolis尼科波利斯
	f.尼科波利斯Nikopolis.SetParent(f)
	
	f.塞尔迪卡Serdica = serdica.CSerdica塞尔迪卡
	f.塞尔迪卡Serdica.SetParent(f)
	
	f.特尔诺沃Tyrnovo = tyrnovo.CTyrnovo特尔诺沃
	f.特尔诺沃Tyrnovo.SetParent(f)
	
}
