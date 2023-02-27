package bhutan

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/bhutan/bumthang"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/bhutan/paro"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BhutanDuke interface {
    feud.Duke
    CBumthang布姆唐() 	bumthang.BumthangCounty
    CParo帕罗() 	paro.ParoCounty
}

type 不丹BhutanDuke struct {
	feud.BaseDuke
	布姆唐Bumthang 	bumthang.BumthangCounty
	帕罗Paro 	paro.ParoCounty
}

func (d *不丹BhutanDuke) CBumthang布姆唐() bumthang.BumthangCounty {
	return d.布姆唐Bumthang
}
    
func (d *不丹BhutanDuke) CParo帕罗() paro.ParoCounty {
	return d.帕罗Paro
}
    
var DBhutan不丹 BhutanDuke = &不丹BhutanDuke{}

func init() {
	f := DBhutan不丹.(*不丹BhutanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bhutan",
		TitleName: "不丹",
		TitleCode: "d_bhutan",
		Counties:  map[string]feud.County{},
	}

	f.布姆唐Bumthang = bumthang.CBumthang布姆唐
	f.布姆唐Bumthang.SetParent(f)
	
	f.帕罗Paro = paro.CParo帕罗
	f.帕罗Paro.SetParent(f)
	
}
