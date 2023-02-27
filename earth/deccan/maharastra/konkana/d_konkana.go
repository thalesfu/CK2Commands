package konkana

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/konkana/goa"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/konkana/honnore"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/konkana/thana"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KonkanaDuke interface {
    feud.Duke
    CGoa果阿() 	goa.GoaCounty
    CHonnore胡奴梨() 	honnore.HonnoreCounty
    CThana都奴何() 	thana.ThanaCounty
}

type 恭建那KonkanaDuke struct {
	feud.BaseDuke
	果阿Goa 	goa.GoaCounty
	胡奴梨Honnore 	honnore.HonnoreCounty
	都奴何Thana 	thana.ThanaCounty
}

func (d *恭建那KonkanaDuke) CGoa果阿() goa.GoaCounty {
	return d.果阿Goa
}
    
func (d *恭建那KonkanaDuke) CHonnore胡奴梨() honnore.HonnoreCounty {
	return d.胡奴梨Honnore
}
    
func (d *恭建那KonkanaDuke) CThana都奴何() thana.ThanaCounty {
	return d.都奴何Thana
}
    
var DKonkana恭建那 KonkanaDuke = &恭建那KonkanaDuke{}

func init() {
	f := DKonkana恭建那.(*恭建那KonkanaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "konkana",
		TitleName: "恭建那",
		TitleCode: "d_konkana",
		Counties:  map[string]feud.County{},
	}

	f.果阿Goa = goa.CGoa果阿
	f.果阿Goa.SetParent(f)
	
	f.胡奴梨Honnore = honnore.CHonnore胡奴梨
	f.胡奴梨Honnore.SetParent(f)
	
	f.都奴何Thana = thana.CThana都奴何
	f.都奴何Thana.SetParent(f)
	
}
