package murom

import (
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/murom/murom"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MuromDuke interface {
	feud.Duke
	CMurom穆罗姆() murom.MuromCounty
}

type 穆罗姆MuromDuke struct {
	feud.BaseDuke
	穆罗姆Murom murom.MuromCounty
}

func (d *穆罗姆MuromDuke) CMurom穆罗姆() murom.MuromCounty {
	return d.穆罗姆Murom
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

}
