package barkul

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/barkul/aj_bogd"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/barkul/barkul"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BarkulDuke interface {
    feud.Duke
    CAj_bogd阿主孛黑答() 	aj_bogd.Aj_bogdCounty
    CBarkul婆悉厥() 	barkul.BarkulCounty
}

type 婆悉厥BarkulDuke struct {
	feud.BaseDuke
	阿主孛黑答Aj_bogd 	aj_bogd.Aj_bogdCounty
	婆悉厥Barkul 	barkul.BarkulCounty
}

func (d *婆悉厥BarkulDuke) CAj_bogd阿主孛黑答() aj_bogd.Aj_bogdCounty {
	return d.阿主孛黑答Aj_bogd
}
    
func (d *婆悉厥BarkulDuke) CBarkul婆悉厥() barkul.BarkulCounty {
	return d.婆悉厥Barkul
}
    
var DBarkul婆悉厥 BarkulDuke = &婆悉厥BarkulDuke{}

func init() {
	f := DBarkul婆悉厥.(*婆悉厥BarkulDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "barkul",
		TitleName: "婆悉厥",
		TitleCode: "d_barkul",
		Counties:  map[string]feud.County{},
	}

	f.阿主孛黑答Aj_bogd = aj_bogd.CAj_bogd阿主孛黑答
	f.阿主孛黑答Aj_bogd.SetParent(f)
	
	f.婆悉厥Barkul = barkul.CBarkul婆悉厥
	f.婆悉厥Barkul.SetParent(f)
	
}
