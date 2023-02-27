package mesopotamia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/mesopotamia/karin"
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/mesopotamia/mesopotamia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/mesopotamia/taron"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MesopotamiaDuke interface {
    feud.Duke
    CKarin卡林() 	karin.KarinCounty
    CMesopotamia赫利亚特() 	mesopotamia.MesopotamiaCounty
    CTaron塔龙() 	taron.TaronCounty
}

type 美索不达米亚MesopotamiaDuke struct {
	feud.BaseDuke
	卡林Karin 	karin.KarinCounty
	赫利亚特Mesopotamia 	mesopotamia.MesopotamiaCounty
	塔龙Taron 	taron.TaronCounty
}

func (d *美索不达米亚MesopotamiaDuke) CKarin卡林() karin.KarinCounty {
	return d.卡林Karin
}
    
func (d *美索不达米亚MesopotamiaDuke) CMesopotamia赫利亚特() mesopotamia.MesopotamiaCounty {
	return d.赫利亚特Mesopotamia
}
    
func (d *美索不达米亚MesopotamiaDuke) CTaron塔龙() taron.TaronCounty {
	return d.塔龙Taron
}
    
var DMesopotamia美索不达米亚 MesopotamiaDuke = &美索不达米亚MesopotamiaDuke{}

func init() {
	f := DMesopotamia美索不达米亚.(*美索不达米亚MesopotamiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "mesopotamia",
		TitleName: "美索不达米亚",
		TitleCode: "d_mesopotamia",
		Counties:  map[string]feud.County{},
	}

	f.卡林Karin = karin.CKarin卡林
	f.卡林Karin.SetParent(f)
	
	f.赫利亚特Mesopotamia = mesopotamia.CMesopotamia赫利亚特
	f.赫利亚特Mesopotamia.SetParent(f)
	
	f.塔龙Taron = taron.CTaron塔龙
	f.塔龙Taron.SetParent(f)
	
}
