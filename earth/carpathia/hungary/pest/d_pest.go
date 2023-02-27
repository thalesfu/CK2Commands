package pest

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/pest/csanad"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/pest/heves"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/pest/pest"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PestDuke interface {
    feud.Duke
    CCsanad乔纳德() 	csanad.CsanadCounty
    CHeves赫维什() 	heves.HevesCounty
    CPest佩斯() 	pest.PestCounty
}

type 佩斯PestDuke struct {
	feud.BaseDuke
	乔纳德Csanad 	csanad.CsanadCounty
	赫维什Heves 	heves.HevesCounty
	佩斯Pest 	pest.PestCounty
}

func (d *佩斯PestDuke) CCsanad乔纳德() csanad.CsanadCounty {
	return d.乔纳德Csanad
}
    
func (d *佩斯PestDuke) CHeves赫维什() heves.HevesCounty {
	return d.赫维什Heves
}
    
func (d *佩斯PestDuke) CPest佩斯() pest.PestCounty {
	return d.佩斯Pest
}
    
var DPest佩斯 PestDuke = &佩斯PestDuke{}

func init() {
	f := DPest佩斯.(*佩斯PestDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "pest",
		TitleName: "佩斯",
		TitleCode: "d_pest",
		Counties:  map[string]feud.County{},
	}

	f.乔纳德Csanad = csanad.CCsanad乔纳德
	f.乔纳德Csanad.SetParent(f)
	
	f.赫维什Heves = heves.CHeves赫维什
	f.赫维什Heves.SetParent(f)
	
	f.佩斯Pest = pest.CPest佩斯
	f.佩斯Pest.SetParent(f)
	
}
