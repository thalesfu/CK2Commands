package alexandria

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/alexandria/alexandria"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/alexandria/gabiyaha"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/alexandria/gizeh"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/alexandria/kharibta"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlexandriaDuke interface {
    feud.Duke
    CAlexandria亚历山大() 	alexandria.AlexandriaCounty
    CGabiyaha拉希德() 	gabiyaha.GabiyahaCounty
    CGizeh法尤姆() 	gizeh.GizehCounty
    CKharibta海里卜塔() 	kharibta.KharibtaCounty
}

type 亚历山大AlexandriaDuke struct {
	feud.BaseDuke
	亚历山大Alexandria 	alexandria.AlexandriaCounty
	拉希德Gabiyaha 	gabiyaha.GabiyahaCounty
	法尤姆Gizeh 	gizeh.GizehCounty
	海里卜塔Kharibta 	kharibta.KharibtaCounty
}

func (d *亚历山大AlexandriaDuke) CAlexandria亚历山大() alexandria.AlexandriaCounty {
	return d.亚历山大Alexandria
}
    
func (d *亚历山大AlexandriaDuke) CGabiyaha拉希德() gabiyaha.GabiyahaCounty {
	return d.拉希德Gabiyaha
}
    
func (d *亚历山大AlexandriaDuke) CGizeh法尤姆() gizeh.GizehCounty {
	return d.法尤姆Gizeh
}
    
func (d *亚历山大AlexandriaDuke) CKharibta海里卜塔() kharibta.KharibtaCounty {
	return d.海里卜塔Kharibta
}
    
var DAlexandria亚历山大 AlexandriaDuke = &亚历山大AlexandriaDuke{}

func init() {
	f := DAlexandria亚历山大.(*亚历山大AlexandriaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "alexandria",
		TitleName: "亚历山大",
		TitleCode: "d_alexandria",
		Counties:  map[string]feud.County{},
	}

	f.亚历山大Alexandria = alexandria.CAlexandria亚历山大
	f.亚历山大Alexandria.SetParent(f)
	
	f.拉希德Gabiyaha = gabiyaha.CGabiyaha拉希德
	f.拉希德Gabiyaha.SetParent(f)
	
	f.法尤姆Gizeh = gizeh.CGizeh法尤姆
	f.法尤姆Gizeh.SetParent(f)
	
	f.海里卜塔Kharibta = kharibta.CKharibta海里卜塔
	f.海里卜塔Kharibta.SetParent(f)
	
}
