package nobatia

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/nobatia/aydhab"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/nobatia/dotawo"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/nobatia/nobatia"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/nobatia/nubia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NobatiaDuke interface {
    feud.Duke
    CAydhab阿伊迪哈卜() 	aydhab.AydhabCounty
    CDotawo多塔沃() 	dotawo.DotawoCounty
    CNobatia诺巴提亚() 	nobatia.NobatiaCounty
    CNubia努比亚() 	nubia.NubiaCounty
}

type 诺巴提亚NobatiaDuke struct {
	feud.BaseDuke
	阿伊迪哈卜Aydhab 	aydhab.AydhabCounty
	多塔沃Dotawo 	dotawo.DotawoCounty
	诺巴提亚Nobatia 	nobatia.NobatiaCounty
	努比亚Nubia 	nubia.NubiaCounty
}

func (d *诺巴提亚NobatiaDuke) CAydhab阿伊迪哈卜() aydhab.AydhabCounty {
	return d.阿伊迪哈卜Aydhab
}
    
func (d *诺巴提亚NobatiaDuke) CDotawo多塔沃() dotawo.DotawoCounty {
	return d.多塔沃Dotawo
}
    
func (d *诺巴提亚NobatiaDuke) CNobatia诺巴提亚() nobatia.NobatiaCounty {
	return d.诺巴提亚Nobatia
}
    
func (d *诺巴提亚NobatiaDuke) CNubia努比亚() nubia.NubiaCounty {
	return d.努比亚Nubia
}
    
var DNobatia诺巴提亚 NobatiaDuke = &诺巴提亚NobatiaDuke{}

func init() {
	f := DNobatia诺巴提亚.(*诺巴提亚NobatiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "nobatia",
		TitleName: "诺巴提亚",
		TitleCode: "d_nobatia",
		Counties:  map[string]feud.County{},
	}

	f.阿伊迪哈卜Aydhab = aydhab.CAydhab阿伊迪哈卜
	f.阿伊迪哈卜Aydhab.SetParent(f)
	
	f.多塔沃Dotawo = dotawo.CDotawo多塔沃
	f.多塔沃Dotawo.SetParent(f)
	
	f.诺巴提亚Nobatia = nobatia.CNobatia诺巴提亚
	f.诺巴提亚Nobatia.SetParent(f)
	
	f.努比亚Nubia = nubia.CNubia努比亚
	f.努比亚Nubia.SetParent(f)
	
}
