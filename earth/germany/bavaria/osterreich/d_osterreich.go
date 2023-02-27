package osterreich

import (
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/osterreich/freistadt"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/osterreich/krems"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/osterreich/medelike"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/osterreich/osterreich"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/osterreich/steiermark"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/osterreich/traungau"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OsterreichDuke interface {
    feud.Duke
    CFreistadt弗赖施塔特() 	freistadt.FreistadtCounty
    CKrems克雷米斯() 	krems.KremsCounty
    CMedelike梅德利克() 	medelike.MedelikeCounty
    COsterreich维恩() 	osterreich.OsterreichCounty
    CSteiermark施蒂里亚() 	steiermark.SteiermarkCounty
    CTraungau特劳恩高() 	traungau.TraungauCounty
}

type 奥地利OsterreichDuke struct {
	feud.BaseDuke
	弗赖施塔特Freistadt 	freistadt.FreistadtCounty
	克雷米斯Krems 	krems.KremsCounty
	梅德利克Medelike 	medelike.MedelikeCounty
	维恩Osterreich 	osterreich.OsterreichCounty
	施蒂里亚Steiermark 	steiermark.SteiermarkCounty
	特劳恩高Traungau 	traungau.TraungauCounty
}

func (d *奥地利OsterreichDuke) CFreistadt弗赖施塔特() freistadt.FreistadtCounty {
	return d.弗赖施塔特Freistadt
}
    
func (d *奥地利OsterreichDuke) CKrems克雷米斯() krems.KremsCounty {
	return d.克雷米斯Krems
}
    
func (d *奥地利OsterreichDuke) CMedelike梅德利克() medelike.MedelikeCounty {
	return d.梅德利克Medelike
}
    
func (d *奥地利OsterreichDuke) COsterreich维恩() osterreich.OsterreichCounty {
	return d.维恩Osterreich
}
    
func (d *奥地利OsterreichDuke) CSteiermark施蒂里亚() steiermark.SteiermarkCounty {
	return d.施蒂里亚Steiermark
}
    
func (d *奥地利OsterreichDuke) CTraungau特劳恩高() traungau.TraungauCounty {
	return d.特劳恩高Traungau
}
    
var DOsterreich奥地利 OsterreichDuke = &奥地利OsterreichDuke{}

func init() {
	f := DOsterreich奥地利.(*奥地利OsterreichDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "osterreich",
		TitleName: "奥地利",
		TitleCode: "d_osterreich",
		Counties:  map[string]feud.County{},
	}

	f.弗赖施塔特Freistadt = freistadt.CFreistadt弗赖施塔特
	f.弗赖施塔特Freistadt.SetParent(f)
	
	f.克雷米斯Krems = krems.CKrems克雷米斯
	f.克雷米斯Krems.SetParent(f)
	
	f.梅德利克Medelike = medelike.CMedelike梅德利克
	f.梅德利克Medelike.SetParent(f)
	
	f.维恩Osterreich = osterreich.COsterreich维恩
	f.维恩Osterreich.SetParent(f)
	
	f.施蒂里亚Steiermark = steiermark.CSteiermark施蒂里亚
	f.施蒂里亚Steiermark.SetParent(f)
	
	f.特劳恩高Traungau = traungau.CTraungau特劳恩高
	f.特劳恩高Traungau.SetParent(f)
	
}
