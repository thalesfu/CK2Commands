package yaik

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/yaik/magnitaya"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/yaik/or"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/yaik/yaik"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YaikDuke interface {
    feud.Duke
    CMagnitaya马格尼特纳亚() 	magnitaya.MagnitayaCounty
    COr奥尔() 	or.OrCounty
    CYaik亚伊克() 	yaik.YaikCounty
}

type 亚伊克YaikDuke struct {
	feud.BaseDuke
	马格尼特纳亚Magnitaya 	magnitaya.MagnitayaCounty
	奥尔Or 	or.OrCounty
	亚伊克Yaik 	yaik.YaikCounty
}

func (d *亚伊克YaikDuke) CMagnitaya马格尼特纳亚() magnitaya.MagnitayaCounty {
	return d.马格尼特纳亚Magnitaya
}
    
func (d *亚伊克YaikDuke) COr奥尔() or.OrCounty {
	return d.奥尔Or
}
    
func (d *亚伊克YaikDuke) CYaik亚伊克() yaik.YaikCounty {
	return d.亚伊克Yaik
}
    
var DYaik亚伊克 YaikDuke = &亚伊克YaikDuke{}

func init() {
	f := DYaik亚伊克.(*亚伊克YaikDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "yaik",
		TitleName: "亚伊克",
		TitleCode: "d_yaik",
		Counties:  map[string]feud.County{},
	}

	f.马格尼特纳亚Magnitaya = magnitaya.CMagnitaya马格尼特纳亚
	f.马格尼特纳亚Magnitaya.SetParent(f)
	
	f.奥尔Or = or.COr奥尔
	f.奥尔Or.SetParent(f)
	
	f.亚伊克Yaik = yaik.CYaik亚伊克
	f.亚伊克Yaik.SetParent(f)
	
}
