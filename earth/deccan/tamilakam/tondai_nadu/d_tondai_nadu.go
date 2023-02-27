package tondai_nadu

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/tondai_nadu/kanchipuram"
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/tondai_nadu/potapi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Tondai_naduDuke interface {
    feud.Duke
    CKanchipuram建志补罗() 	kanchipuram.KanchipuramCounty
    CPotapi菩多毗() 	potapi.PotapiCounty
}

type 旦榸拿柱Tondai_naduDuke struct {
	feud.BaseDuke
	建志补罗Kanchipuram 	kanchipuram.KanchipuramCounty
	菩多毗Potapi 	potapi.PotapiCounty
}

func (d *旦榸拿柱Tondai_naduDuke) CKanchipuram建志补罗() kanchipuram.KanchipuramCounty {
	return d.建志补罗Kanchipuram
}
    
func (d *旦榸拿柱Tondai_naduDuke) CPotapi菩多毗() potapi.PotapiCounty {
	return d.菩多毗Potapi
}
    
var DTondai_nadu旦榸拿柱 Tondai_naduDuke = &旦榸拿柱Tondai_naduDuke{}

func init() {
	f := DTondai_nadu旦榸拿柱.(*旦榸拿柱Tondai_naduDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tondai_nadu",
		TitleName: "旦榸拿柱",
		TitleCode: "d_tondai_nadu",
		Counties:  map[string]feud.County{},
	}

	f.建志补罗Kanchipuram = kanchipuram.CKanchipuram建志补罗
	f.建志补罗Kanchipuram.SetParent(f)
	
	f.菩多毗Potapi = potapi.CPotapi菩多毗
	f.菩多毗Potapi.SetParent(f)
	
}
