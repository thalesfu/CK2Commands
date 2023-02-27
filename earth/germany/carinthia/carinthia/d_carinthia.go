package carinthia

import (
	"github.com/thalesfu/CK2Commands/earth/germany/carinthia/carinthia/friesach"
	"github.com/thalesfu/CK2Commands/earth/germany/carinthia/carinthia/karnten"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CarinthiaDuke interface {
    feud.Duke
    CFriesach弗里萨赫() 	friesach.FriesachCounty
    CKarnten菲拉赫() 	karnten.KarntenCounty
}

type 卡林西亚CarinthiaDuke struct {
	feud.BaseDuke
	弗里萨赫Friesach 	friesach.FriesachCounty
	菲拉赫Karnten 	karnten.KarntenCounty
}

func (d *卡林西亚CarinthiaDuke) CFriesach弗里萨赫() friesach.FriesachCounty {
	return d.弗里萨赫Friesach
}
    
func (d *卡林西亚CarinthiaDuke) CKarnten菲拉赫() karnten.KarntenCounty {
	return d.菲拉赫Karnten
}
    
var DCarinthia卡林西亚 CarinthiaDuke = &卡林西亚CarinthiaDuke{}

func init() {
	f := DCarinthia卡林西亚.(*卡林西亚CarinthiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "carinthia",
		TitleName: "卡林西亚",
		TitleCode: "d_carinthia",
		Counties:  map[string]feud.County{},
	}

	f.弗里萨赫Friesach = friesach.CFriesach弗里萨赫
	f.弗里萨赫Friesach.SetParent(f)
	
	f.菲拉赫Karnten = karnten.CKarnten菲拉赫
	f.菲拉赫Karnten.SetParent(f)
	
}
