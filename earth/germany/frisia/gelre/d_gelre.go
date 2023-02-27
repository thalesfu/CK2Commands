package gelre

import (
	"github.com/thalesfu/CK2Commands/earth/germany/frisia/gelre/frisia"
	"github.com/thalesfu/CK2Commands/earth/germany/frisia/gelre/gelre"
	"github.com/thalesfu/CK2Commands/earth/germany/frisia/gelre/ostfriesland"
	"github.com/thalesfu/CK2Commands/earth/germany/frisia/gelre/overijssel"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GelreDuke interface {
    feud.Duke
    CFrisia弗里西亚() 	frisia.FrisiaCounty
    CGelre海尔雷() 	gelre.GelreCounty
    COstfriesland东弗里斯兰() 	ostfriesland.OstfrieslandCounty
    COverijssel上艾瑟尔() 	overijssel.OverijsselCounty
}

type 海尔雷GelreDuke struct {
	feud.BaseDuke
	弗里西亚Frisia 	frisia.FrisiaCounty
	海尔雷Gelre 	gelre.GelreCounty
	东弗里斯兰Ostfriesland 	ostfriesland.OstfrieslandCounty
	上艾瑟尔Overijssel 	overijssel.OverijsselCounty
}

func (d *海尔雷GelreDuke) CFrisia弗里西亚() frisia.FrisiaCounty {
	return d.弗里西亚Frisia
}
    
func (d *海尔雷GelreDuke) CGelre海尔雷() gelre.GelreCounty {
	return d.海尔雷Gelre
}
    
func (d *海尔雷GelreDuke) COstfriesland东弗里斯兰() ostfriesland.OstfrieslandCounty {
	return d.东弗里斯兰Ostfriesland
}
    
func (d *海尔雷GelreDuke) COverijssel上艾瑟尔() overijssel.OverijsselCounty {
	return d.上艾瑟尔Overijssel
}
    
var DGelre海尔雷 GelreDuke = &海尔雷GelreDuke{}

func init() {
	f := DGelre海尔雷.(*海尔雷GelreDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gelre",
		TitleName: "海尔雷",
		TitleCode: "d_gelre",
		Counties:  map[string]feud.County{},
	}

	f.弗里西亚Frisia = frisia.CFrisia弗里西亚
	f.弗里西亚Frisia.SetParent(f)
	
	f.海尔雷Gelre = gelre.CGelre海尔雷
	f.海尔雷Gelre.SetParent(f)
	
	f.东弗里斯兰Ostfriesland = ostfriesland.COstfriesland东弗里斯兰
	f.东弗里斯兰Ostfriesland.SetParent(f)
	
	f.上艾瑟尔Overijssel = overijssel.COverijssel上艾瑟尔
	f.上艾瑟尔Overijssel.SetParent(f)
	
}
