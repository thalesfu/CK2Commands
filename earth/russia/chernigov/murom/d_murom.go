package murom

import (
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/murom/murom"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/murom/orekhovo_zouievo"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MuromDuke interface {
    feud.Duke
    CMurom穆罗姆() 	murom.MuromCounty
    COrekhovo_zouievo奥列霍沃_祖耶沃() 	orekhovo_zouievo.Orekhovo_zouievoCounty
}

type 穆罗姆MuromDuke struct {
	feud.BaseDuke
	穆罗姆Murom 	murom.MuromCounty
	奥列霍沃_祖耶沃Orekhovo_zouievo 	orekhovo_zouievo.Orekhovo_zouievoCounty
}

func (d *穆罗姆MuromDuke) CMurom穆罗姆() murom.MuromCounty {
	return d.穆罗姆Murom
}
    
func (d *穆罗姆MuromDuke) COrekhovo_zouievo奥列霍沃_祖耶沃() orekhovo_zouievo.Orekhovo_zouievoCounty {
	return d.奥列霍沃_祖耶沃Orekhovo_zouievo
}
    
var DMurom穆罗姆 MuromDuke = &穆罗姆MuromDuke{}

func init() {
	f := DMurom穆罗姆.(*穆罗姆MuromDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "murom",
		TitleName: "穆罗姆",
		TitleCode: "d_murom",
		Counties:  map[string]feud.County{},
	}

	f.穆罗姆Murom = murom.CMurom穆罗姆
	f.穆罗姆Murom.SetParent(f)
	
	f.奥列霍沃_祖耶沃Orekhovo_zouievo = orekhovo_zouievo.COrekhovo_zouievo奥列霍沃_祖耶沃
	f.奥列霍沃_祖耶沃Orekhovo_zouievo.SetParent(f)
	
}
