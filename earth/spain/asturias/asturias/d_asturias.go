package asturias

import (
	"github.com/thalesfu/CK2Commands/earth/spain/asturias/asturias/astorga"
	"github.com/thalesfu/CK2Commands/earth/spain/asturias/asturias/asturias_de_oviedo"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AsturiasDuke interface {
    feud.Duke
    CAstorga阿斯托加() 	astorga.AstorgaCounty
    CAsturias_de_oviedo阿斯图里亚斯德奥维耶多() 	asturias_de_oviedo.Asturias_de_oviedoCounty
}

type 阿斯图里亚斯AsturiasDuke struct {
	feud.BaseDuke
	阿斯托加Astorga 	astorga.AstorgaCounty
	阿斯图里亚斯德奥维耶多Asturias_de_oviedo 	asturias_de_oviedo.Asturias_de_oviedoCounty
}

func (d *阿斯图里亚斯AsturiasDuke) CAstorga阿斯托加() astorga.AstorgaCounty {
	return d.阿斯托加Astorga
}
    
func (d *阿斯图里亚斯AsturiasDuke) CAsturias_de_oviedo阿斯图里亚斯德奥维耶多() asturias_de_oviedo.Asturias_de_oviedoCounty {
	return d.阿斯图里亚斯德奥维耶多Asturias_de_oviedo
}
    
var DAsturias阿斯图里亚斯 AsturiasDuke = &阿斯图里亚斯AsturiasDuke{}

func init() {
	f := DAsturias阿斯图里亚斯.(*阿斯图里亚斯AsturiasDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "asturias",
		TitleName: "阿斯图里亚斯",
		TitleCode: "d_asturias",
		Counties:  map[string]feud.County{},
	}

	f.阿斯托加Astorga = astorga.CAstorga阿斯托加
	f.阿斯托加Astorga.SetParent(f)
	
	f.阿斯图里亚斯德奥维耶多Asturias_de_oviedo = asturias_de_oviedo.CAsturias_de_oviedo阿斯图里亚斯德奥维耶多
	f.阿斯图里亚斯德奥维耶多Asturias_de_oviedo.SetParent(f)
	
}
