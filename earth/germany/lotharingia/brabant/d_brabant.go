package brabant

import (
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/brabant/brabant"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/brabant/breda"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/brabant/hainaut"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BrabantDuke interface {
    feud.Duke
    CBrabant布拉班特() 	brabant.BrabantCounty
    CBreda布雷达() 	breda.BredaCounty
    CHainaut埃诺() 	hainaut.HainautCounty
}

type 布拉班特BrabantDuke struct {
	feud.BaseDuke
	布拉班特Brabant 	brabant.BrabantCounty
	布雷达Breda 	breda.BredaCounty
	埃诺Hainaut 	hainaut.HainautCounty
}

func (d *布拉班特BrabantDuke) CBrabant布拉班特() brabant.BrabantCounty {
	return d.布拉班特Brabant
}
    
func (d *布拉班特BrabantDuke) CBreda布雷达() breda.BredaCounty {
	return d.布雷达Breda
}
    
func (d *布拉班特BrabantDuke) CHainaut埃诺() hainaut.HainautCounty {
	return d.埃诺Hainaut
}
    
var DBrabant布拉班特 BrabantDuke = &布拉班特BrabantDuke{}

func init() {
	f := DBrabant布拉班特.(*布拉班特BrabantDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "brabant",
		TitleName: "布拉班特",
		TitleCode: "d_brabant",
		Counties:  map[string]feud.County{},
	}

	f.布拉班特Brabant = brabant.CBrabant布拉班特
	f.布拉班特Brabant.SetParent(f)
	
	f.布雷达Breda = breda.CBreda布雷达
	f.布雷达Breda.SetParent(f)
	
	f.埃诺Hainaut = hainaut.CHainaut埃诺
	f.埃诺Hainaut.SetParent(f)
	
}
