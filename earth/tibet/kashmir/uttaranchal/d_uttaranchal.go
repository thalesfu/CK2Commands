package uttaranchal

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/kashmir/uttaranchal/garhwal"
	"github.com/thalesfu/CK2Commands/earth/tibet/kashmir/uttaranchal/kangra"
	"github.com/thalesfu/CK2Commands/earth/tibet/kashmir/uttaranchal/kurmanchal"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UttaranchalDuke interface {
    feud.Duke
    CGarhwal真日() 	garhwal.GarhwalCounty
    CKangra建伽罗() 	kangra.KangraCounty
    CKurmanchal俱利曼遮罗() 	kurmanchal.KurmanchalCounty
}

type 北安遮罗UttaranchalDuke struct {
	feud.BaseDuke
	真日Garhwal 	garhwal.GarhwalCounty
	建伽罗Kangra 	kangra.KangraCounty
	俱利曼遮罗Kurmanchal 	kurmanchal.KurmanchalCounty
}

func (d *北安遮罗UttaranchalDuke) CGarhwal真日() garhwal.GarhwalCounty {
	return d.真日Garhwal
}
    
func (d *北安遮罗UttaranchalDuke) CKangra建伽罗() kangra.KangraCounty {
	return d.建伽罗Kangra
}
    
func (d *北安遮罗UttaranchalDuke) CKurmanchal俱利曼遮罗() kurmanchal.KurmanchalCounty {
	return d.俱利曼遮罗Kurmanchal
}
    
var DUttaranchal北安遮罗 UttaranchalDuke = &北安遮罗UttaranchalDuke{}

func init() {
	f := DUttaranchal北安遮罗.(*北安遮罗UttaranchalDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "uttaranchal",
		TitleName: "北安遮罗",
		TitleCode: "d_uttaranchal",
		Counties:  map[string]feud.County{},
	}

	f.真日Garhwal = garhwal.CGarhwal真日
	f.真日Garhwal.SetParent(f)
	
	f.建伽罗Kangra = kangra.CKangra建伽罗
	f.建伽罗Kangra.SetParent(f)
	
	f.俱利曼遮罗Kurmanchal = kurmanchal.CKurmanchal俱利曼遮罗
	f.俱利曼遮罗Kurmanchal.SetParent(f)
	
}
