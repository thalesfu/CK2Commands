package flanders

import (
	"github.com/thalesfu/CK2Commands/earth/france/france/flanders/artois"
	"github.com/thalesfu/CK2Commands/earth/france/france/flanders/boulogne"
	"github.com/thalesfu/CK2Commands/earth/france/france/flanders/brugge"
	"github.com/thalesfu/CK2Commands/earth/france/france/flanders/gent"
	"github.com/thalesfu/CK2Commands/earth/france/france/flanders/guines"
	"github.com/thalesfu/CK2Commands/earth/france/france/flanders/yperen"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FlandersDuke interface {
    feud.Duke
    CArtois阿图瓦() 	artois.ArtoisCounty
    CBoulogne布洛涅() 	boulogne.BoulogneCounty
    CBrugge布鲁日() 	brugge.BruggeCounty
    CGent根特() 	gent.GentCounty
    CGuines吉讷() 	guines.GuinesCounty
    CYperen伊普尔() 	yperen.YperenCounty
}

type 佛兰德斯FlandersDuke struct {
	feud.BaseDuke
	阿图瓦Artois 	artois.ArtoisCounty
	布洛涅Boulogne 	boulogne.BoulogneCounty
	布鲁日Brugge 	brugge.BruggeCounty
	根特Gent 	gent.GentCounty
	吉讷Guines 	guines.GuinesCounty
	伊普尔Yperen 	yperen.YperenCounty
}

func (d *佛兰德斯FlandersDuke) CArtois阿图瓦() artois.ArtoisCounty {
	return d.阿图瓦Artois
}
    
func (d *佛兰德斯FlandersDuke) CBoulogne布洛涅() boulogne.BoulogneCounty {
	return d.布洛涅Boulogne
}
    
func (d *佛兰德斯FlandersDuke) CBrugge布鲁日() brugge.BruggeCounty {
	return d.布鲁日Brugge
}
    
func (d *佛兰德斯FlandersDuke) CGent根特() gent.GentCounty {
	return d.根特Gent
}
    
func (d *佛兰德斯FlandersDuke) CGuines吉讷() guines.GuinesCounty {
	return d.吉讷Guines
}
    
func (d *佛兰德斯FlandersDuke) CYperen伊普尔() yperen.YperenCounty {
	return d.伊普尔Yperen
}
    
var DFlanders佛兰德斯 FlandersDuke = &佛兰德斯FlandersDuke{}

func init() {
	f := DFlanders佛兰德斯.(*佛兰德斯FlandersDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "flanders",
		TitleName: "佛兰德斯",
		TitleCode: "d_flanders",
		Counties:  map[string]feud.County{},
	}

	f.阿图瓦Artois = artois.CArtois阿图瓦
	f.阿图瓦Artois.SetParent(f)
	
	f.布洛涅Boulogne = boulogne.CBoulogne布洛涅
	f.布洛涅Boulogne.SetParent(f)
	
	f.布鲁日Brugge = brugge.CBrugge布鲁日
	f.布鲁日Brugge.SetParent(f)
	
	f.根特Gent = gent.CGent根特
	f.根特Gent.SetParent(f)
	
	f.吉讷Guines = guines.CGuines吉讷
	f.吉讷Guines.SetParent(f)
	
	f.伊普尔Yperen = yperen.CYperen伊普尔
	f.伊普尔Yperen.SetParent(f)
	
}
