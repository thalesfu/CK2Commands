package leon

import (
	"github.com/thalesfu/CK2Commands/earth/spain/asturias/leon/leon"
	"github.com/thalesfu/CK2Commands/earth/spain/asturias/leon/palencia"
	"github.com/thalesfu/CK2Commands/earth/spain/asturias/leon/salamanca"
	"github.com/thalesfu/CK2Commands/earth/spain/asturias/leon/zamora"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LeonDuke interface {
    feud.Duke
    CLeon莱昂() 	leon.LeonCounty
    CPalencia帕伦西亚() 	palencia.PalenciaCounty
    CSalamanca萨拉曼卡() 	salamanca.SalamancaCounty
    CZamora萨莫拉() 	zamora.ZamoraCounty
}

type 莱昂LeonDuke struct {
	feud.BaseDuke
	莱昂Leon 	leon.LeonCounty
	帕伦西亚Palencia 	palencia.PalenciaCounty
	萨拉曼卡Salamanca 	salamanca.SalamancaCounty
	萨莫拉Zamora 	zamora.ZamoraCounty
}

func (d *莱昂LeonDuke) CLeon莱昂() leon.LeonCounty {
	return d.莱昂Leon
}
    
func (d *莱昂LeonDuke) CPalencia帕伦西亚() palencia.PalenciaCounty {
	return d.帕伦西亚Palencia
}
    
func (d *莱昂LeonDuke) CSalamanca萨拉曼卡() salamanca.SalamancaCounty {
	return d.萨拉曼卡Salamanca
}
    
func (d *莱昂LeonDuke) CZamora萨莫拉() zamora.ZamoraCounty {
	return d.萨莫拉Zamora
}
    
var DLeon莱昂 LeonDuke = &莱昂LeonDuke{}

func init() {
	f := DLeon莱昂.(*莱昂LeonDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "leon",
		TitleName: "莱昂",
		TitleCode: "d_leon",
		Counties:  map[string]feud.County{},
	}

	f.莱昂Leon = leon.CLeon莱昂
	f.莱昂Leon.SetParent(f)
	
	f.帕伦西亚Palencia = palencia.CPalencia帕伦西亚
	f.帕伦西亚Palencia.SetParent(f)
	
	f.萨拉曼卡Salamanca = salamanca.CSalamanca萨拉曼卡
	f.萨拉曼卡Salamanca.SetParent(f)
	
	f.萨莫拉Zamora = zamora.CZamora萨莫拉
	f.萨莫拉Zamora.SetParent(f)
	
}
