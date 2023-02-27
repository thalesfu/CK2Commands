package abydos

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/abydos/abydos"
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/abydos/kyzikos"
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/abydos/lesbos"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AbydosDuke interface {
    feud.Duke
    CAbydos阿卑多斯() 	abydos.AbydosCounty
    CKyzikos基齐库斯() 	kyzikos.KyzikosCounty
    CLesbos莱斯沃斯() 	lesbos.LesbosCounty
}

type 阿卑多斯AbydosDuke struct {
	feud.BaseDuke
	阿卑多斯Abydos 	abydos.AbydosCounty
	基齐库斯Kyzikos 	kyzikos.KyzikosCounty
	莱斯沃斯Lesbos 	lesbos.LesbosCounty
}

func (d *阿卑多斯AbydosDuke) CAbydos阿卑多斯() abydos.AbydosCounty {
	return d.阿卑多斯Abydos
}
    
func (d *阿卑多斯AbydosDuke) CKyzikos基齐库斯() kyzikos.KyzikosCounty {
	return d.基齐库斯Kyzikos
}
    
func (d *阿卑多斯AbydosDuke) CLesbos莱斯沃斯() lesbos.LesbosCounty {
	return d.莱斯沃斯Lesbos
}
    
var DAbydos阿卑多斯 AbydosDuke = &阿卑多斯AbydosDuke{}

func init() {
	f := DAbydos阿卑多斯.(*阿卑多斯AbydosDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "abydos",
		TitleName: "阿卑多斯",
		TitleCode: "d_abydos",
		Counties:  map[string]feud.County{},
	}

	f.阿卑多斯Abydos = abydos.CAbydos阿卑多斯
	f.阿卑多斯Abydos.SetParent(f)
	
	f.基齐库斯Kyzikos = kyzikos.CKyzikos基齐库斯
	f.基齐库斯Kyzikos.SetParent(f)
	
	f.莱斯沃斯Lesbos = lesbos.CLesbos莱斯沃斯
	f.莱斯沃斯Lesbos.SetParent(f)
	
}
