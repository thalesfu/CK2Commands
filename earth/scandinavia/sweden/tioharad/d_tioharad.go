package tioharad

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/tioharad/finnveden"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/tioharad/njudung"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/tioharad/smaland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TioharadDuke interface {
    feud.Duke
    CFinnveden芬韦登() 	finnveden.FinnvedenCounty
    CNjudung纽东() 	njudung.NjudungCounty
    CSmaland韦伦德() 	smaland.SmalandCounty
}

type 韦伦德TioharadDuke struct {
	feud.BaseDuke
	芬韦登Finnveden 	finnveden.FinnvedenCounty
	纽东Njudung 	njudung.NjudungCounty
	韦伦德Smaland 	smaland.SmalandCounty
}

func (d *韦伦德TioharadDuke) CFinnveden芬韦登() finnveden.FinnvedenCounty {
	return d.芬韦登Finnveden
}
    
func (d *韦伦德TioharadDuke) CNjudung纽东() njudung.NjudungCounty {
	return d.纽东Njudung
}
    
func (d *韦伦德TioharadDuke) CSmaland韦伦德() smaland.SmalandCounty {
	return d.韦伦德Smaland
}
    
var DTioharad韦伦德 TioharadDuke = &韦伦德TioharadDuke{}

func init() {
	f := DTioharad韦伦德.(*韦伦德TioharadDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tioharad",
		TitleName: "韦伦德",
		TitleCode: "d_tioharad",
		Counties:  map[string]feud.County{},
	}

	f.芬韦登Finnveden = finnveden.CFinnveden芬韦登
	f.芬韦登Finnveden.SetParent(f)
	
	f.纽东Njudung = njudung.CNjudung纽东
	f.纽东Njudung.SetParent(f)
	
	f.韦伦德Smaland = smaland.CSmaland韦伦德
	f.韦伦德Smaland.SetParent(f)
	
}
