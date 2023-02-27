package dalmatia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/dalmatia/imotski"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/dalmatia/zachlumia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DalmatiaDuke interface {
    feud.Duke
    CImotski伊莫茨基() 	imotski.ImotskiCounty
    CZachlumia扎库卢米亚() 	zachlumia.ZachlumiaCounty
}

type 胡姆DalmatiaDuke struct {
	feud.BaseDuke
	伊莫茨基Imotski 	imotski.ImotskiCounty
	扎库卢米亚Zachlumia 	zachlumia.ZachlumiaCounty
}

func (d *胡姆DalmatiaDuke) CImotski伊莫茨基() imotski.ImotskiCounty {
	return d.伊莫茨基Imotski
}
    
func (d *胡姆DalmatiaDuke) CZachlumia扎库卢米亚() zachlumia.ZachlumiaCounty {
	return d.扎库卢米亚Zachlumia
}
    
var DDalmatia胡姆 DalmatiaDuke = &胡姆DalmatiaDuke{}

func init() {
	f := DDalmatia胡姆.(*胡姆DalmatiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "dalmatia",
		TitleName: "胡姆",
		TitleCode: "d_dalmatia",
		Counties:  map[string]feud.County{},
	}

	f.伊莫茨基Imotski = imotski.CImotski伊莫茨基
	f.伊莫茨基Imotski.SetParent(f)
	
	f.扎库卢米亚Zachlumia = zachlumia.CZachlumia扎库卢米亚
	f.扎库卢米亚Zachlumia.SetParent(f)
	
}
