package strymon

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/strymon/astibus"
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/strymon/strymon"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type StrymonDuke interface {
    feud.Duke
    CAstibus阿斯提博斯() 	astibus.AstibusCounty
    CStrymon斯特里蒙() 	strymon.StrymonCounty
}

type 斯特里蒙StrymonDuke struct {
	feud.BaseDuke
	阿斯提博斯Astibus 	astibus.AstibusCounty
	斯特里蒙Strymon 	strymon.StrymonCounty
}

func (d *斯特里蒙StrymonDuke) CAstibus阿斯提博斯() astibus.AstibusCounty {
	return d.阿斯提博斯Astibus
}
    
func (d *斯特里蒙StrymonDuke) CStrymon斯特里蒙() strymon.StrymonCounty {
	return d.斯特里蒙Strymon
}
    
var DStrymon斯特里蒙 StrymonDuke = &斯特里蒙StrymonDuke{}

func init() {
	f := DStrymon斯特里蒙.(*斯特里蒙StrymonDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "strymon",
		TitleName: "斯特里蒙",
		TitleCode: "d_strymon",
		Counties:  map[string]feud.County{},
	}

	f.阿斯提博斯Astibus = astibus.CAstibus阿斯提博斯
	f.阿斯提博斯Astibus.SetParent(f)
	
	f.斯特里蒙Strymon = strymon.CStrymon斯特里蒙
	f.斯特里蒙Strymon.SetParent(f)
	
}
