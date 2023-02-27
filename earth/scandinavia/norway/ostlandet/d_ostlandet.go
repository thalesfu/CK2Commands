package ostlandet

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/ostlandet/akershus"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/ostlandet/vestfold"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/ostlandet/viken"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OstlandetDuke interface {
    feud.Duke
    CAkershus阿克什胡斯() 	akershus.AkershusCounty
    CVestfold西福尔() 	vestfold.VestfoldCounty
    CViken博胡斯伦() 	viken.VikenCounty
}

type 维肯OstlandetDuke struct {
	feud.BaseDuke
	阿克什胡斯Akershus 	akershus.AkershusCounty
	西福尔Vestfold 	vestfold.VestfoldCounty
	博胡斯伦Viken 	viken.VikenCounty
}

func (d *维肯OstlandetDuke) CAkershus阿克什胡斯() akershus.AkershusCounty {
	return d.阿克什胡斯Akershus
}
    
func (d *维肯OstlandetDuke) CVestfold西福尔() vestfold.VestfoldCounty {
	return d.西福尔Vestfold
}
    
func (d *维肯OstlandetDuke) CViken博胡斯伦() viken.VikenCounty {
	return d.博胡斯伦Viken
}
    
var DOstlandet维肯 OstlandetDuke = &维肯OstlandetDuke{}

func init() {
	f := DOstlandet维肯.(*维肯OstlandetDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ostlandet",
		TitleName: "维肯",
		TitleCode: "d_ostlandet",
		Counties:  map[string]feud.County{},
	}

	f.阿克什胡斯Akershus = akershus.CAkershus阿克什胡斯
	f.阿克什胡斯Akershus.SetParent(f)
	
	f.西福尔Vestfold = vestfold.CVestfold西福尔
	f.西福尔Vestfold.SetParent(f)
	
	f.博胡斯伦Viken = viken.CViken博胡斯伦
	f.博胡斯伦Viken.SetParent(f)
	
}
