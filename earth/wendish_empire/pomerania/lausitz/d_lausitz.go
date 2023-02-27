package lausitz

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/lausitz/anhalt"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/lausitz/lausitz"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/lausitz/plauen"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LausitzDuke interface {
    feud.Duke
    CAnhalt采尔布斯特() 	anhalt.AnhaltCounty
    CLausitz劳西茨() 	lausitz.LausitzCounty
    CPlauen梅泽堡() 	plauen.PlauenCounty
}

type 劳西茨LausitzDuke struct {
	feud.BaseDuke
	采尔布斯特Anhalt 	anhalt.AnhaltCounty
	劳西茨Lausitz 	lausitz.LausitzCounty
	梅泽堡Plauen 	plauen.PlauenCounty
}

func (d *劳西茨LausitzDuke) CAnhalt采尔布斯特() anhalt.AnhaltCounty {
	return d.采尔布斯特Anhalt
}
    
func (d *劳西茨LausitzDuke) CLausitz劳西茨() lausitz.LausitzCounty {
	return d.劳西茨Lausitz
}
    
func (d *劳西茨LausitzDuke) CPlauen梅泽堡() plauen.PlauenCounty {
	return d.梅泽堡Plauen
}
    
var DLausitz劳西茨 LausitzDuke = &劳西茨LausitzDuke{}

func init() {
	f := DLausitz劳西茨.(*劳西茨LausitzDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "lausitz",
		TitleName: "劳西茨",
		TitleCode: "d_lausitz",
		Counties:  map[string]feud.County{},
	}

	f.采尔布斯特Anhalt = anhalt.CAnhalt采尔布斯特
	f.采尔布斯特Anhalt.SetParent(f)
	
	f.劳西茨Lausitz = lausitz.CLausitz劳西茨
	f.劳西茨Lausitz.SetParent(f)
	
	f.梅泽堡Plauen = plauen.CPlauen梅泽堡
	f.梅泽堡Plauen.SetParent(f)
	
}
