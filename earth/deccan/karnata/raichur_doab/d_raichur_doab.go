package raichur_doab

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/raichur_doab/banavasi"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/raichur_doab/idatarainadu"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/raichur_doab/kudalasangama"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata/raichur_doab/vatapi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Raichur_doabDuke interface {
    feud.Duke
    CBanavasi巴南巴西() 	banavasi.BanavasiCounty
    CIdatarainadu伊陀多来拿都() 	idatarainadu.IdatarainaduCounty
    CKudalasangama俱荼罗僧伽摩() 	kudalasangama.KudalasangamaCounty
    CVatapi伐达比() 	vatapi.VatapiCounty
}

type 罗耶注卢河间地Raichur_doabDuke struct {
	feud.BaseDuke
	巴南巴西Banavasi 	banavasi.BanavasiCounty
	伊陀多来拿都Idatarainadu 	idatarainadu.IdatarainaduCounty
	俱荼罗僧伽摩Kudalasangama 	kudalasangama.KudalasangamaCounty
	伐达比Vatapi 	vatapi.VatapiCounty
}

func (d *罗耶注卢河间地Raichur_doabDuke) CBanavasi巴南巴西() banavasi.BanavasiCounty {
	return d.巴南巴西Banavasi
}
    
func (d *罗耶注卢河间地Raichur_doabDuke) CIdatarainadu伊陀多来拿都() idatarainadu.IdatarainaduCounty {
	return d.伊陀多来拿都Idatarainadu
}
    
func (d *罗耶注卢河间地Raichur_doabDuke) CKudalasangama俱荼罗僧伽摩() kudalasangama.KudalasangamaCounty {
	return d.俱荼罗僧伽摩Kudalasangama
}
    
func (d *罗耶注卢河间地Raichur_doabDuke) CVatapi伐达比() vatapi.VatapiCounty {
	return d.伐达比Vatapi
}
    
var DRaichur_doab罗耶注卢河间地 Raichur_doabDuke = &罗耶注卢河间地Raichur_doabDuke{}

func init() {
	f := DRaichur_doab罗耶注卢河间地.(*罗耶注卢河间地Raichur_doabDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "raichur_doab",
		TitleName: "罗耶注卢河间地",
		TitleCode: "d_raichur_doab",
		Counties:  map[string]feud.County{},
	}

	f.巴南巴西Banavasi = banavasi.CBanavasi巴南巴西
	f.巴南巴西Banavasi.SetParent(f)
	
	f.伊陀多来拿都Idatarainadu = idatarainadu.CIdatarainadu伊陀多来拿都
	f.伊陀多来拿都Idatarainadu.SetParent(f)
	
	f.俱荼罗僧伽摩Kudalasangama = kudalasangama.CKudalasangama俱荼罗僧伽摩
	f.俱荼罗僧伽摩Kudalasangama.SetParent(f)
	
	f.伐达比Vatapi = vatapi.CVatapi伐达比
	f.伐达比Vatapi.SetParent(f)
	
}
