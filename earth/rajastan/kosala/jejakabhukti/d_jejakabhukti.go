package jejakabhukti

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/jejakabhukti/damoh"
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/jejakabhukti/gurgi"
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/jejakabhukti/kalanjara"
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/jejakabhukti/mahoba"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JejakabhuktiDuke interface {
    feud.Duke
    CDamoh佗牟() 	damoh.DamohCounty
    CGurgi古卢耆() 	gurgi.GurgiCounty
    CKalanjara迦兰阇罗() 	kalanjara.KalanjaraCounty
    CMahoba摩呼婆() 	mahoba.MahobaCounty
}

type 掷枳陀JejakabhuktiDuke struct {
	feud.BaseDuke
	佗牟Damoh 	damoh.DamohCounty
	古卢耆Gurgi 	gurgi.GurgiCounty
	迦兰阇罗Kalanjara 	kalanjara.KalanjaraCounty
	摩呼婆Mahoba 	mahoba.MahobaCounty
}

func (d *掷枳陀JejakabhuktiDuke) CDamoh佗牟() damoh.DamohCounty {
	return d.佗牟Damoh
}
    
func (d *掷枳陀JejakabhuktiDuke) CGurgi古卢耆() gurgi.GurgiCounty {
	return d.古卢耆Gurgi
}
    
func (d *掷枳陀JejakabhuktiDuke) CKalanjara迦兰阇罗() kalanjara.KalanjaraCounty {
	return d.迦兰阇罗Kalanjara
}
    
func (d *掷枳陀JejakabhuktiDuke) CMahoba摩呼婆() mahoba.MahobaCounty {
	return d.摩呼婆Mahoba
}
    
var DJejakabhukti掷枳陀 JejakabhuktiDuke = &掷枳陀JejakabhuktiDuke{}

func init() {
	f := DJejakabhukti掷枳陀.(*掷枳陀JejakabhuktiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "jejakabhukti",
		TitleName: "掷枳陀",
		TitleCode: "d_jejakabhukti",
		Counties:  map[string]feud.County{},
	}

	f.佗牟Damoh = damoh.CDamoh佗牟
	f.佗牟Damoh.SetParent(f)
	
	f.古卢耆Gurgi = gurgi.CGurgi古卢耆
	f.古卢耆Gurgi.SetParent(f)
	
	f.迦兰阇罗Kalanjara = kalanjara.CKalanjara迦兰阇罗
	f.迦兰阇罗Kalanjara.SetParent(f)
	
	f.摩呼婆Mahoba = mahoba.CMahoba摩呼婆
	f.摩呼婆Mahoba.SetParent(f)
	
}
