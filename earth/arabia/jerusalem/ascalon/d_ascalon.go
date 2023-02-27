package ascalon

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/ascalon/ascalon"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/ascalon/beersheb"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/ascalon/darum"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/ascalon/jaffa"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AscalonDuke interface {
    feud.Duke
    CAscalon亚实基伦() 	ascalon.AscalonCounty
    CBeersheb贝尔谢巴() 	beersheb.BeershebCounty
    CDarum达鲁姆() 	darum.DarumCounty
    CJaffa雅法() 	jaffa.JaffaCounty
}

type 亚实基伦AscalonDuke struct {
	feud.BaseDuke
	亚实基伦Ascalon 	ascalon.AscalonCounty
	贝尔谢巴Beersheb 	beersheb.BeershebCounty
	达鲁姆Darum 	darum.DarumCounty
	雅法Jaffa 	jaffa.JaffaCounty
}

func (d *亚实基伦AscalonDuke) CAscalon亚实基伦() ascalon.AscalonCounty {
	return d.亚实基伦Ascalon
}
    
func (d *亚实基伦AscalonDuke) CBeersheb贝尔谢巴() beersheb.BeershebCounty {
	return d.贝尔谢巴Beersheb
}
    
func (d *亚实基伦AscalonDuke) CDarum达鲁姆() darum.DarumCounty {
	return d.达鲁姆Darum
}
    
func (d *亚实基伦AscalonDuke) CJaffa雅法() jaffa.JaffaCounty {
	return d.雅法Jaffa
}
    
var DAscalon亚实基伦 AscalonDuke = &亚实基伦AscalonDuke{}

func init() {
	f := DAscalon亚实基伦.(*亚实基伦AscalonDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ascalon",
		TitleName: "亚实基伦",
		TitleCode: "d_ascalon",
		Counties:  map[string]feud.County{},
	}

	f.亚实基伦Ascalon = ascalon.CAscalon亚实基伦
	f.亚实基伦Ascalon.SetParent(f)
	
	f.贝尔谢巴Beersheb = beersheb.CBeersheb贝尔谢巴
	f.贝尔谢巴Beersheb.SetParent(f)
	
	f.达鲁姆Darum = darum.CDarum达鲁姆
	f.达鲁姆Darum.SetParent(f)
	
	f.雅法Jaffa = jaffa.CJaffa雅法
	f.雅法Jaffa.SetParent(f)
	
}
