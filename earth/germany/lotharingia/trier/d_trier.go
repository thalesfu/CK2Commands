package trier

import (
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/trier/metz"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/trier/saargau"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/trier/trier"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrierDuke interface {
    feud.Duke
    CMetz梅斯() 	metz.MetzCounty
    CSaargau萨尔布吕肯() 	saargau.SaargauCounty
    CTrier特里尔() 	trier.TrierCounty
}

type 特里尔TrierDuke struct {
	feud.BaseDuke
	梅斯Metz 	metz.MetzCounty
	萨尔布吕肯Saargau 	saargau.SaargauCounty
	特里尔Trier 	trier.TrierCounty
}

func (d *特里尔TrierDuke) CMetz梅斯() metz.MetzCounty {
	return d.梅斯Metz
}
    
func (d *特里尔TrierDuke) CSaargau萨尔布吕肯() saargau.SaargauCounty {
	return d.萨尔布吕肯Saargau
}
    
func (d *特里尔TrierDuke) CTrier特里尔() trier.TrierCounty {
	return d.特里尔Trier
}
    
var DTrier特里尔 TrierDuke = &特里尔TrierDuke{}

func init() {
	f := DTrier特里尔.(*特里尔TrierDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "trier",
		TitleName: "特里尔",
		TitleCode: "d_trier",
		Counties:  map[string]feud.County{},
	}

	f.梅斯Metz = metz.CMetz梅斯
	f.梅斯Metz.SetParent(f)
	
	f.萨尔布吕肯Saargau = saargau.CSaargau萨尔布吕肯
	f.萨尔布吕肯Saargau.SetParent(f)
	
	f.特里尔Trier = trier.CTrier特里尔
	f.特里尔Trier.SetParent(f)
	
}
