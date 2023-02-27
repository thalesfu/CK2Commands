package atyrau

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/atyrau/guryev"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/atyrau/ryn"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AtyrauDuke interface {
    feud.Duke
    CGuryev阿特劳() 	guryev.GuryevCounty
    CRyn雷恩() 	ryn.RynCounty
}

type 阿特劳AtyrauDuke struct {
	feud.BaseDuke
	阿特劳Guryev 	guryev.GuryevCounty
	雷恩Ryn 	ryn.RynCounty
}

func (d *阿特劳AtyrauDuke) CGuryev阿特劳() guryev.GuryevCounty {
	return d.阿特劳Guryev
}
    
func (d *阿特劳AtyrauDuke) CRyn雷恩() ryn.RynCounty {
	return d.雷恩Ryn
}
    
var DAtyrau阿特劳 AtyrauDuke = &阿特劳AtyrauDuke{}

func init() {
	f := DAtyrau阿特劳.(*阿特劳AtyrauDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "atyrau",
		TitleName: "阿特劳",
		TitleCode: "d_atyrau",
		Counties:  map[string]feud.County{},
	}

	f.阿特劳Guryev = guryev.CGuryev阿特劳
	f.阿特劳Guryev.SetParent(f)
	
	f.雷恩Ryn = ryn.CRyn雷恩
	f.雷恩Ryn.SetParent(f)
	
}
