package genoa

import (
	"github.com/thalesfu/CK2Commands/earth/italy/italy/genoa/genoa"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/genoa/noli"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GenoaDuke interface {
    feud.Duke
    CGenoa热那亚() 	genoa.GenoaCounty
    CNoli诺利() 	noli.NoliCounty
}

type 热那亚GenoaDuke struct {
	feud.BaseDuke
	热那亚Genoa 	genoa.GenoaCounty
	诺利Noli 	noli.NoliCounty
}

func (d *热那亚GenoaDuke) CGenoa热那亚() genoa.GenoaCounty {
	return d.热那亚Genoa
}
    
func (d *热那亚GenoaDuke) CNoli诺利() noli.NoliCounty {
	return d.诺利Noli
}
    
var DGenoa热那亚 GenoaDuke = &热那亚GenoaDuke{}

func init() {
	f := DGenoa热那亚.(*热那亚GenoaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "genoa",
		TitleName: "热那亚",
		TitleCode: "d_genoa",
		Counties:  map[string]feud.County{},
	}

	f.热那亚Genoa = genoa.CGenoa热那亚
	f.热那亚Genoa.SetParent(f)
	
	f.诺利Noli = noli.CNoli诺利
	f.诺利Noli.SetParent(f)
	
}
