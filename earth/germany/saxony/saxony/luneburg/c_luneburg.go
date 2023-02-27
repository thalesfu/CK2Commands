package luneburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LuneburgCounty interface {
    feud.County
    BBardowick巴多维克() 	feud.Barony
    BEvern埃弗恩() 	feud.Barony
    BGifhorn吉夫霍恩() 	feud.Barony
    BLudersburg吕德斯堡() 	feud.Barony
    BLuneburg吕讷堡() 	feud.Barony
    BReppenstedt雷彭施泰特() 	feud.Barony
    BThomasburg托马斯堡() 	feud.Barony
    BUelzen于尔岑() 	feud.Barony
}

type 吕讷堡LuneburgCounty struct {
	feud.BaseCounty
	巴多维克Bardowick 	feud.Barony
	埃弗恩Evern 	feud.Barony
	吉夫霍恩Gifhorn 	feud.Barony
	吕德斯堡Ludersburg 	feud.Barony
	吕讷堡Luneburg 	feud.Barony
	雷彭施泰特Reppenstedt 	feud.Barony
	托马斯堡Thomasburg 	feud.Barony
	于尔岑Uelzen 	feud.Barony
}

func (c *吕讷堡LuneburgCounty) BBardowick巴多维克() feud.Barony {
	return c.巴多维克Bardowick
}
    
func (c *吕讷堡LuneburgCounty) BEvern埃弗恩() feud.Barony {
	return c.埃弗恩Evern
}
    
func (c *吕讷堡LuneburgCounty) BGifhorn吉夫霍恩() feud.Barony {
	return c.吉夫霍恩Gifhorn
}
    
func (c *吕讷堡LuneburgCounty) BLudersburg吕德斯堡() feud.Barony {
	return c.吕德斯堡Ludersburg
}
    
func (c *吕讷堡LuneburgCounty) BLuneburg吕讷堡() feud.Barony {
	return c.吕讷堡Luneburg
}
    
func (c *吕讷堡LuneburgCounty) BReppenstedt雷彭施泰特() feud.Barony {
	return c.雷彭施泰特Reppenstedt
}
    
func (c *吕讷堡LuneburgCounty) BThomasburg托马斯堡() feud.Barony {
	return c.托马斯堡Thomasburg
}
    
func (c *吕讷堡LuneburgCounty) BUelzen于尔岑() feud.Barony {
	return c.于尔岑Uelzen
}
    
var CLuneburg吕讷堡 LuneburgCounty = &吕讷堡LuneburgCounty{}

func init() {
	f := CLuneburg吕讷堡.(*吕讷堡LuneburgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "258",
		Title:     "luneburg",
		TitleName: "吕讷堡",
		TitleCode: "c_luneburg",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴多维克Bardowick = BBardowick巴多维克
	f.巴多维克Bardowick.SetParent(f)
	
	f.埃弗恩Evern = BEvern埃弗恩
	f.埃弗恩Evern.SetParent(f)
	
	f.吉夫霍恩Gifhorn = BGifhorn吉夫霍恩
	f.吉夫霍恩Gifhorn.SetParent(f)
	
	f.吕德斯堡Ludersburg = BLudersburg吕德斯堡
	f.吕德斯堡Ludersburg.SetParent(f)
	
	f.吕讷堡Luneburg = BLuneburg吕讷堡
	f.吕讷堡Luneburg.SetParent(f)
	
	f.雷彭施泰特Reppenstedt = BReppenstedt雷彭施泰特
	f.雷彭施泰特Reppenstedt.SetParent(f)
	
	f.托马斯堡Thomasburg = BThomasburg托马斯堡
	f.托马斯堡Thomasburg.SetParent(f)
	
	f.于尔岑Uelzen = BUelzen于尔岑
	f.于尔岑Uelzen.SetParent(f)
	
}
