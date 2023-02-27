package latgale

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/latgale/jersika"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/latgale/selija"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/latgale/talava"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/latgale/west_dvina"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LatgaleDuke interface {
    feud.Duke
    CJersika耶尔西卡() 	jersika.JersikaCounty
    CSelija塞丽亚() 	selija.SelijaCounty
    CTalava塔拉瓦() 	talava.TalavaCounty
    CWest_dvina西德维纳() 	west_dvina.West_dvinaCounty
}

type 拉特加尔LatgaleDuke struct {
	feud.BaseDuke
	耶尔西卡Jersika 	jersika.JersikaCounty
	塞丽亚Selija 	selija.SelijaCounty
	塔拉瓦Talava 	talava.TalavaCounty
	西德维纳West_dvina 	west_dvina.West_dvinaCounty
}

func (d *拉特加尔LatgaleDuke) CJersika耶尔西卡() jersika.JersikaCounty {
	return d.耶尔西卡Jersika
}
    
func (d *拉特加尔LatgaleDuke) CSelija塞丽亚() selija.SelijaCounty {
	return d.塞丽亚Selija
}
    
func (d *拉特加尔LatgaleDuke) CTalava塔拉瓦() talava.TalavaCounty {
	return d.塔拉瓦Talava
}
    
func (d *拉特加尔LatgaleDuke) CWest_dvina西德维纳() west_dvina.West_dvinaCounty {
	return d.西德维纳West_dvina
}
    
var DLatgale拉特加尔 LatgaleDuke = &拉特加尔LatgaleDuke{}

func init() {
	f := DLatgale拉特加尔.(*拉特加尔LatgaleDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "latgale",
		TitleName: "拉特加尔",
		TitleCode: "d_latgale",
		Counties:  map[string]feud.County{},
	}

	f.耶尔西卡Jersika = jersika.CJersika耶尔西卡
	f.耶尔西卡Jersika.SetParent(f)
	
	f.塞丽亚Selija = selija.CSelija塞丽亚
	f.塞丽亚Selija.SetParent(f)
	
	f.塔拉瓦Talava = talava.CTalava塔拉瓦
	f.塔拉瓦Talava.SetParent(f)
	
	f.西德维纳West_dvina = west_dvina.CWest_dvina西德维纳
	f.西德维纳West_dvina.SetParent(f)
	
}
