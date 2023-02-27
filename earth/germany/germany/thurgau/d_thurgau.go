package thurgau

import (
	"github.com/thalesfu/CK2Commands/earth/germany/germany/thurgau/st_gallen"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/thurgau/zurichgau"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThurgauDuke interface {
    feud.Duke
    CSt_gallen图尔高() 	st_gallen.St_gallenCounty
    CZurichgau苏黎世() 	zurichgau.ZurichgauCounty
}

type 图尔高ThurgauDuke struct {
	feud.BaseDuke
	图尔高St_gallen 	st_gallen.St_gallenCounty
	苏黎世Zurichgau 	zurichgau.ZurichgauCounty
}

func (d *图尔高ThurgauDuke) CSt_gallen图尔高() st_gallen.St_gallenCounty {
	return d.图尔高St_gallen
}
    
func (d *图尔高ThurgauDuke) CZurichgau苏黎世() zurichgau.ZurichgauCounty {
	return d.苏黎世Zurichgau
}
    
var DThurgau图尔高 ThurgauDuke = &图尔高ThurgauDuke{}

func init() {
	f := DThurgau图尔高.(*图尔高ThurgauDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "thurgau",
		TitleName: "图尔高",
		TitleCode: "d_thurgau",
		Counties:  map[string]feud.County{},
	}

	f.图尔高St_gallen = st_gallen.CSt_gallen图尔高
	f.图尔高St_gallen.SetParent(f)
	
	f.苏黎世Zurichgau = zurichgau.CZurichgau苏黎世
	f.苏黎世Zurichgau.SetParent(f)
	
}
