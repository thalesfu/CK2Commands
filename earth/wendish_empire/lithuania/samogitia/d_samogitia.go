package samogitia

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/samogitia/memel"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/samogitia/scalovia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SamogitiaDuke interface {
    feud.Duke
    CMemel梅梅尔() 	memel.MemelCounty
    CScalovia斯卡洛维亚() 	scalovia.ScaloviaCounty
}

type 斯卡尔维亚SamogitiaDuke struct {
	feud.BaseDuke
	梅梅尔Memel 	memel.MemelCounty
	斯卡洛维亚Scalovia 	scalovia.ScaloviaCounty
}

func (d *斯卡尔维亚SamogitiaDuke) CMemel梅梅尔() memel.MemelCounty {
	return d.梅梅尔Memel
}
    
func (d *斯卡尔维亚SamogitiaDuke) CScalovia斯卡洛维亚() scalovia.ScaloviaCounty {
	return d.斯卡洛维亚Scalovia
}
    
var DSamogitia斯卡尔维亚 SamogitiaDuke = &斯卡尔维亚SamogitiaDuke{}

func init() {
	f := DSamogitia斯卡尔维亚.(*斯卡尔维亚SamogitiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "samogitia",
		TitleName: "斯卡尔维亚",
		TitleCode: "d_samogitia",
		Counties:  map[string]feud.County{},
	}

	f.梅梅尔Memel = memel.CMemel梅梅尔
	f.梅梅尔Memel.SetParent(f)
	
	f.斯卡洛维亚Scalovia = scalovia.CScalovia斯卡洛维亚
	f.斯卡洛维亚Scalovia.SetParent(f)
	
}
