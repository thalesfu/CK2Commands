package trigarta

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/trigarta/gurjaratra"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/trigarta/sakala"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/trigarta/trigarta"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrigartaDuke interface {
    feud.Duke
    CGurjaratra瞿折罗多罗() 	gurjaratra.GurjaratraCounty
    CSakala奢羯罗() 	sakala.SakalaCounty
    CTrigarta三穴() 	trigarta.TrigartaCounty
}

type 三穴TrigartaDuke struct {
	feud.BaseDuke
	瞿折罗多罗Gurjaratra 	gurjaratra.GurjaratraCounty
	奢羯罗Sakala 	sakala.SakalaCounty
	三穴Trigarta 	trigarta.TrigartaCounty
}

func (d *三穴TrigartaDuke) CGurjaratra瞿折罗多罗() gurjaratra.GurjaratraCounty {
	return d.瞿折罗多罗Gurjaratra
}
    
func (d *三穴TrigartaDuke) CSakala奢羯罗() sakala.SakalaCounty {
	return d.奢羯罗Sakala
}
    
func (d *三穴TrigartaDuke) CTrigarta三穴() trigarta.TrigartaCounty {
	return d.三穴Trigarta
}
    
var DTrigarta三穴 TrigartaDuke = &三穴TrigartaDuke{}

func init() {
	f := DTrigarta三穴.(*三穴TrigartaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "trigarta",
		TitleName: "三穴",
		TitleCode: "d_trigarta",
		Counties:  map[string]feud.County{},
	}

	f.瞿折罗多罗Gurjaratra = gurjaratra.CGurjaratra瞿折罗多罗
	f.瞿折罗多罗Gurjaratra.SetParent(f)
	
	f.奢羯罗Sakala = sakala.CSakala奢羯罗
	f.奢羯罗Sakala.SetParent(f)
	
	f.三穴Trigarta = trigarta.CTrigarta三穴
	f.三穴Trigarta.SetParent(f)
	
}
