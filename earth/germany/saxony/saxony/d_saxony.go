package saxony

import (
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/saxony/altmark"
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/saxony/braunschweig"
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/saxony/luneburg"
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/saxony/magdeburg"
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/saxony/verden"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaxonyDuke interface {
    feud.Duke
    CAltmark韦尔本() 	altmark.AltmarkCounty
    CBraunschweig布伦瑞克() 	braunschweig.BraunschweigCounty
    CLuneburg吕讷堡() 	luneburg.LuneburgCounty
    CMagdeburg马格德堡() 	magdeburg.MagdeburgCounty
    CVerden费尔登() 	verden.VerdenCounty
}

type 萨克森SaxonyDuke struct {
	feud.BaseDuke
	韦尔本Altmark 	altmark.AltmarkCounty
	布伦瑞克Braunschweig 	braunschweig.BraunschweigCounty
	吕讷堡Luneburg 	luneburg.LuneburgCounty
	马格德堡Magdeburg 	magdeburg.MagdeburgCounty
	费尔登Verden 	verden.VerdenCounty
}

func (d *萨克森SaxonyDuke) CAltmark韦尔本() altmark.AltmarkCounty {
	return d.韦尔本Altmark
}
    
func (d *萨克森SaxonyDuke) CBraunschweig布伦瑞克() braunschweig.BraunschweigCounty {
	return d.布伦瑞克Braunschweig
}
    
func (d *萨克森SaxonyDuke) CLuneburg吕讷堡() luneburg.LuneburgCounty {
	return d.吕讷堡Luneburg
}
    
func (d *萨克森SaxonyDuke) CMagdeburg马格德堡() magdeburg.MagdeburgCounty {
	return d.马格德堡Magdeburg
}
    
func (d *萨克森SaxonyDuke) CVerden费尔登() verden.VerdenCounty {
	return d.费尔登Verden
}
    
var DSaxony萨克森 SaxonyDuke = &萨克森SaxonyDuke{}

func init() {
	f := DSaxony萨克森.(*萨克森SaxonyDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "saxony",
		TitleName: "萨克森",
		TitleCode: "d_saxony",
		Counties:  map[string]feud.County{},
	}

	f.韦尔本Altmark = altmark.CAltmark韦尔本
	f.韦尔本Altmark.SetParent(f)
	
	f.布伦瑞克Braunschweig = braunschweig.CBraunschweig布伦瑞克
	f.布伦瑞克Braunschweig.SetParent(f)
	
	f.吕讷堡Luneburg = luneburg.CLuneburg吕讷堡
	f.吕讷堡Luneburg.SetParent(f)
	
	f.马格德堡Magdeburg = magdeburg.CMagdeburg马格德堡
	f.马格德堡Magdeburg.SetParent(f)
	
	f.费尔登Verden = verden.CVerden费尔登
	f.费尔登Verden.SetParent(f)
	
}
