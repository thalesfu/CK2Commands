package greater_poland

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/greater_poland/gnieznienskie"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/greater_poland/kaliskie"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/greater_poland/lubusz"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/greater_poland/poznanskie"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Greater_polandDuke interface {
    feud.Duke
    CGnieznienskie纳克沃() 	gnieznienskie.GnieznienskieCounty
    CKaliskie卡利什() 	kaliskie.KaliskieCounty
    CLubusz卢布什() 	lubusz.LubuszCounty
    CPoznanskie波兹南() 	poznanskie.PoznanskieCounty
}

type 大波兰Greater_polandDuke struct {
	feud.BaseDuke
	纳克沃Gnieznienskie 	gnieznienskie.GnieznienskieCounty
	卡利什Kaliskie 	kaliskie.KaliskieCounty
	卢布什Lubusz 	lubusz.LubuszCounty
	波兹南Poznanskie 	poznanskie.PoznanskieCounty
}

func (d *大波兰Greater_polandDuke) CGnieznienskie纳克沃() gnieznienskie.GnieznienskieCounty {
	return d.纳克沃Gnieznienskie
}
    
func (d *大波兰Greater_polandDuke) CKaliskie卡利什() kaliskie.KaliskieCounty {
	return d.卡利什Kaliskie
}
    
func (d *大波兰Greater_polandDuke) CLubusz卢布什() lubusz.LubuszCounty {
	return d.卢布什Lubusz
}
    
func (d *大波兰Greater_polandDuke) CPoznanskie波兹南() poznanskie.PoznanskieCounty {
	return d.波兹南Poznanskie
}
    
var DGreater_poland大波兰 Greater_polandDuke = &大波兰Greater_polandDuke{}

func init() {
	f := DGreater_poland大波兰.(*大波兰Greater_polandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "greater_poland",
		TitleName: "大波兰",
		TitleCode: "d_greater_poland",
		Counties:  map[string]feud.County{},
	}

	f.纳克沃Gnieznienskie = gnieznienskie.CGnieznienskie纳克沃
	f.纳克沃Gnieznienskie.SetParent(f)
	
	f.卡利什Kaliskie = kaliskie.CKaliskie卡利什
	f.卡利什Kaliskie.SetParent(f)
	
	f.卢布什Lubusz = lubusz.CLubusz卢布什
	f.卢布什Lubusz.SetParent(f)
	
	f.波兹南Poznanskie = poznanskie.CPoznanskie波兹南
	f.波兹南Poznanskie.SetParent(f)
	
}
