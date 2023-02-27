package cairo

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/cairo/akhmim"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/cairo/cairo"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/cairo/kolzum"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/cairo/sarqihya"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CairoDuke interface {
    feud.Duke
    CAkhmim阿赫明() 	akhmim.AkhmimCounty
    CCairo开罗() 	cairo.CairoCounty
    CKolzum科勒祖姆() 	kolzum.KolzumCounty
    CSarqihya比勒拜斯() 	sarqihya.SarqihyaCounty
}

type 开罗CairoDuke struct {
	feud.BaseDuke
	阿赫明Akhmim 	akhmim.AkhmimCounty
	开罗Cairo 	cairo.CairoCounty
	科勒祖姆Kolzum 	kolzum.KolzumCounty
	比勒拜斯Sarqihya 	sarqihya.SarqihyaCounty
}

func (d *开罗CairoDuke) CAkhmim阿赫明() akhmim.AkhmimCounty {
	return d.阿赫明Akhmim
}
    
func (d *开罗CairoDuke) CCairo开罗() cairo.CairoCounty {
	return d.开罗Cairo
}
    
func (d *开罗CairoDuke) CKolzum科勒祖姆() kolzum.KolzumCounty {
	return d.科勒祖姆Kolzum
}
    
func (d *开罗CairoDuke) CSarqihya比勒拜斯() sarqihya.SarqihyaCounty {
	return d.比勒拜斯Sarqihya
}
    
var DCairo开罗 CairoDuke = &开罗CairoDuke{}

func init() {
	f := DCairo开罗.(*开罗CairoDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "cairo",
		TitleName: "开罗",
		TitleCode: "d_cairo",
		Counties:  map[string]feud.County{},
	}

	f.阿赫明Akhmim = akhmim.CAkhmim阿赫明
	f.阿赫明Akhmim.SetParent(f)
	
	f.开罗Cairo = cairo.CCairo开罗
	f.开罗Cairo.SetParent(f)
	
	f.科勒祖姆Kolzum = kolzum.CKolzum科勒祖姆
	f.科勒祖姆Kolzum.SetParent(f)
	
	f.比勒拜斯Sarqihya = sarqihya.CSarqihya比勒拜斯
	f.比勒拜斯Sarqihya.SetParent(f)
	
}
