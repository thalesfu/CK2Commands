package nyingchi

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/kham/nyingchi/mainling"
	"github.com/thalesfu/CK2Commands/earth/tibet/kham/nyingchi/medog"
	"github.com/thalesfu/CK2Commands/earth/tibet/kham/nyingchi/nyingchi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NyingchiDuke interface {
    feud.Duke
    CMainling米林() 	mainling.MainlingCounty
    CMedog墨脱() 	medog.MedogCounty
    CNyingchi尼池() 	nyingchi.NyingchiCounty
}

type 尼池NyingchiDuke struct {
	feud.BaseDuke
	米林Mainling 	mainling.MainlingCounty
	墨脱Medog 	medog.MedogCounty
	尼池Nyingchi 	nyingchi.NyingchiCounty
}

func (d *尼池NyingchiDuke) CMainling米林() mainling.MainlingCounty {
	return d.米林Mainling
}
    
func (d *尼池NyingchiDuke) CMedog墨脱() medog.MedogCounty {
	return d.墨脱Medog
}
    
func (d *尼池NyingchiDuke) CNyingchi尼池() nyingchi.NyingchiCounty {
	return d.尼池Nyingchi
}
    
var DNyingchi尼池 NyingchiDuke = &尼池NyingchiDuke{}

func init() {
	f := DNyingchi尼池.(*尼池NyingchiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "nyingchi",
		TitleName: "尼池",
		TitleCode: "d_nyingchi",
		Counties:  map[string]feud.County{},
	}

	f.米林Mainling = mainling.CMainling米林
	f.米林Mainling.SetParent(f)
	
	f.墨脱Medog = medog.CMedog墨脱
	f.墨脱Medog.SetParent(f)
	
	f.尼池Nyingchi = nyingchi.CNyingchi尼池
	f.尼池Nyingchi.SetParent(f)
	
}
