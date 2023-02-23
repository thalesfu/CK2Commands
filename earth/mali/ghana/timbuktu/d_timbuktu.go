package timbuktu

import (
	"github.com/thalesfu/CK2Commands/earth/mali/ghana/timbuktu/araouane"
	"github.com/thalesfu/CK2Commands/earth/mali/ghana/timbuktu/oualata"
	"github.com/thalesfu/CK2Commands/earth/mali/ghana/timbuktu/timbuktu"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TimbuktuDuke interface {
	feud.Duke
	CAraouane阿拉万() araouane.AraouaneCounty
	COualata瓦拉塔() oualata.OualataCounty
	CTimbuktu廷巴克图() timbuktu.TimbuktuCounty
}

type 廷巴克图TimbuktuDuke struct {
	feud.BaseDuke
	阿拉万Araouane  araouane.AraouaneCounty
	瓦拉塔Oualata   oualata.OualataCounty
	廷巴克图Timbuktu timbuktu.TimbuktuCounty
}

func (d *廷巴克图TimbuktuDuke) CAraouane阿拉万() araouane.AraouaneCounty {
	return d.阿拉万Araouane
}

func (d *廷巴克图TimbuktuDuke) COualata瓦拉塔() oualata.OualataCounty {
	return d.瓦拉塔Oualata
}

func (d *廷巴克图TimbuktuDuke) CTimbuktu廷巴克图() timbuktu.TimbuktuCounty {
	return d.廷巴克图Timbuktu
}

var DTimbuktu廷巴克图 TimbuktuDuke = &廷巴克图TimbuktuDuke{}

func init() {
	f := DTimbuktu廷巴克图.(*廷巴克图TimbuktuDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "timbuktu",
		TitleName: "廷巴克图",
		TitleCode: "d_timbuktu",
		Counties:  map[string]feud.County{},
	}

	f.阿拉万Araouane = araouane.CAraouane阿拉万
	f.阿拉万Araouane.SetParent(f)

	f.瓦拉塔Oualata = oualata.COualata瓦拉塔
	f.瓦拉塔Oualata.SetParent(f)

	f.廷巴克图Timbuktu = timbuktu.CTimbuktu廷巴克图
	f.廷巴克图Timbuktu.SetParent(f)

}
