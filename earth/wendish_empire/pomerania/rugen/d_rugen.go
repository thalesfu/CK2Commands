package rugen

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/rugen/rostock"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/rugen/rugen"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/rugen/werle"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RugenDuke interface {
    feud.Duke
    CRostock罗斯托克() 	rostock.RostockCounty
    CRugen吕根() 	rugen.RugenCounty
    CWerle韦尔勒() 	werle.WerleCounty
}

type 吕根RugenDuke struct {
	feud.BaseDuke
	罗斯托克Rostock 	rostock.RostockCounty
	吕根Rugen 	rugen.RugenCounty
	韦尔勒Werle 	werle.WerleCounty
}

func (d *吕根RugenDuke) CRostock罗斯托克() rostock.RostockCounty {
	return d.罗斯托克Rostock
}
    
func (d *吕根RugenDuke) CRugen吕根() rugen.RugenCounty {
	return d.吕根Rugen
}
    
func (d *吕根RugenDuke) CWerle韦尔勒() werle.WerleCounty {
	return d.韦尔勒Werle
}
    
var DRugen吕根 RugenDuke = &吕根RugenDuke{}

func init() {
	f := DRugen吕根.(*吕根RugenDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "rugen",
		TitleName: "吕根",
		TitleCode: "d_rugen",
		Counties:  map[string]feud.County{},
	}

	f.罗斯托克Rostock = rostock.CRostock罗斯托克
	f.罗斯托克Rostock.SetParent(f)
	
	f.吕根Rugen = rugen.CRugen吕根
	f.吕根Rugen.SetParent(f)
	
	f.韦尔勒Werle = werle.CWerle韦尔勒
	f.韦尔勒Werle.SetParent(f)
	
}
