package meissen

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/meissen/gorlitz"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/meissen/meissen"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MeissenDuke interface {
    feud.Duke
    CGorlitz格尔利茨() 	gorlitz.GorlitzCounty
    CMeissen迈森() 	meissen.MeissenCounty
}

type 迈森MeissenDuke struct {
	feud.BaseDuke
	格尔利茨Gorlitz 	gorlitz.GorlitzCounty
	迈森Meissen 	meissen.MeissenCounty
}

func (d *迈森MeissenDuke) CGorlitz格尔利茨() gorlitz.GorlitzCounty {
	return d.格尔利茨Gorlitz
}
    
func (d *迈森MeissenDuke) CMeissen迈森() meissen.MeissenCounty {
	return d.迈森Meissen
}
    
var DMeissen迈森 MeissenDuke = &迈森MeissenDuke{}

func init() {
	f := DMeissen迈森.(*迈森MeissenDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "meissen",
		TitleName: "迈森",
		TitleCode: "d_meissen",
		Counties:  map[string]feud.County{},
	}

	f.格尔利茨Gorlitz = gorlitz.CGorlitz格尔利茨
	f.格尔利茨Gorlitz.SetParent(f)
	
	f.迈森Meissen = meissen.CMeissen迈森
	f.迈森Meissen.SetParent(f)
	
}
