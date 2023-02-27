package ostergotland

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/ostergotland/kinda"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/ostergotland/ostergotland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/ostergotland/vista"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OstergotlandDuke interface {
    feud.Duke
    CKinda欣达() 	kinda.KindaCounty
    COstergotland东约特兰() 	ostergotland.OstergotlandCounty
    CVista维斯塔() 	vista.VistaCounty
}

type 东约特兰OstergotlandDuke struct {
	feud.BaseDuke
	欣达Kinda 	kinda.KindaCounty
	东约特兰Ostergotland 	ostergotland.OstergotlandCounty
	维斯塔Vista 	vista.VistaCounty
}

func (d *东约特兰OstergotlandDuke) CKinda欣达() kinda.KindaCounty {
	return d.欣达Kinda
}
    
func (d *东约特兰OstergotlandDuke) COstergotland东约特兰() ostergotland.OstergotlandCounty {
	return d.东约特兰Ostergotland
}
    
func (d *东约特兰OstergotlandDuke) CVista维斯塔() vista.VistaCounty {
	return d.维斯塔Vista
}
    
var DOstergotland东约特兰 OstergotlandDuke = &东约特兰OstergotlandDuke{}

func init() {
	f := DOstergotland东约特兰.(*东约特兰OstergotlandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ostergotland",
		TitleName: "东约特兰",
		TitleCode: "d_ostergotland",
		Counties:  map[string]feud.County{},
	}

	f.欣达Kinda = kinda.CKinda欣达
	f.欣达Kinda.SetParent(f)
	
	f.东约特兰Ostergotland = ostergotland.COstergotland东约特兰
	f.东约特兰Ostergotland.SetParent(f)
	
	f.维斯塔Vista = vista.CVista维斯塔
	f.维斯塔Vista.SetParent(f)
	
}
