package wag

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/wag/lakomelza"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/wag/wag"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WagDuke interface {
    feud.Duke
    CLakomelza拉科美尔扎() 	lakomelza.LakomelzaCounty
    CWag瓦格() 	wag.WagCounty
}

type 瓦格WagDuke struct {
	feud.BaseDuke
	拉科美尔扎Lakomelza 	lakomelza.LakomelzaCounty
	瓦格Wag 	wag.WagCounty
}

func (d *瓦格WagDuke) CLakomelza拉科美尔扎() lakomelza.LakomelzaCounty {
	return d.拉科美尔扎Lakomelza
}
    
func (d *瓦格WagDuke) CWag瓦格() wag.WagCounty {
	return d.瓦格Wag
}
    
var DWag瓦格 WagDuke = &瓦格WagDuke{}

func init() {
	f := DWag瓦格.(*瓦格WagDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "wag",
		TitleName: "瓦格",
		TitleCode: "d_wag",
		Counties:  map[string]feud.County{},
	}

	f.拉科美尔扎Lakomelza = lakomelza.CLakomelza拉科美尔扎
	f.拉科美尔扎Lakomelza.SetParent(f)
	
	f.瓦格Wag = wag.CWag瓦格
	f.瓦格Wag.SetParent(f)
	
}
