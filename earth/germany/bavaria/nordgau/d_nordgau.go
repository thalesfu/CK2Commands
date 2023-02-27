package nordgau

import (
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/nordgau/eichstadt"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/nordgau/nurnberg"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NordgauDuke interface {
    feud.Duke
    CEichstadt艾希施泰特() 	eichstadt.EichstadtCounty
    CNurnberg纽伦堡() 	nurnberg.NurnbergCounty
}

type 诺德高NordgauDuke struct {
	feud.BaseDuke
	艾希施泰特Eichstadt 	eichstadt.EichstadtCounty
	纽伦堡Nurnberg 	nurnberg.NurnbergCounty
}

func (d *诺德高NordgauDuke) CEichstadt艾希施泰特() eichstadt.EichstadtCounty {
	return d.艾希施泰特Eichstadt
}
    
func (d *诺德高NordgauDuke) CNurnberg纽伦堡() nurnberg.NurnbergCounty {
	return d.纽伦堡Nurnberg
}
    
var DNordgau诺德高 NordgauDuke = &诺德高NordgauDuke{}

func init() {
	f := DNordgau诺德高.(*诺德高NordgauDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "nordgau",
		TitleName: "诺德高",
		TitleCode: "d_nordgau",
		Counties:  map[string]feud.County{},
	}

	f.艾希施泰特Eichstadt = eichstadt.CEichstadt艾希施泰特
	f.艾希施泰特Eichstadt.SetParent(f)
	
	f.纽伦堡Nurnberg = nurnberg.CNurnberg纽伦堡
	f.纽伦堡Nurnberg.SetParent(f)
	
}
