package calabria

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/calabria/catanzaro"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/calabria/consenza"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/calabria/reggio"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CalabriaDuke interface {
    feud.Duke
    CCatanzaro卡坦扎罗() 	catanzaro.CatanzaroCounty
    CConsenza科森扎() 	consenza.ConsenzaCounty
    CReggio雷焦() 	reggio.ReggioCounty
}

type 卡拉布里亚CalabriaDuke struct {
	feud.BaseDuke
	卡坦扎罗Catanzaro 	catanzaro.CatanzaroCounty
	科森扎Consenza 	consenza.ConsenzaCounty
	雷焦Reggio 	reggio.ReggioCounty
}

func (d *卡拉布里亚CalabriaDuke) CCatanzaro卡坦扎罗() catanzaro.CatanzaroCounty {
	return d.卡坦扎罗Catanzaro
}
    
func (d *卡拉布里亚CalabriaDuke) CConsenza科森扎() consenza.ConsenzaCounty {
	return d.科森扎Consenza
}
    
func (d *卡拉布里亚CalabriaDuke) CReggio雷焦() reggio.ReggioCounty {
	return d.雷焦Reggio
}
    
var DCalabria卡拉布里亚 CalabriaDuke = &卡拉布里亚CalabriaDuke{}

func init() {
	f := DCalabria卡拉布里亚.(*卡拉布里亚CalabriaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "calabria",
		TitleName: "卡拉布里亚",
		TitleCode: "d_calabria",
		Counties:  map[string]feud.County{},
	}

	f.卡坦扎罗Catanzaro = catanzaro.CCatanzaro卡坦扎罗
	f.卡坦扎罗Catanzaro.SetParent(f)
	
	f.科森扎Consenza = consenza.CConsenza科森扎
	f.科森扎Consenza.SetParent(f)
	
	f.雷焦Reggio = reggio.CReggio雷焦
	f.雷焦Reggio.SetParent(f)
	
}
