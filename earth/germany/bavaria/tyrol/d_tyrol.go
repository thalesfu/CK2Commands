package tyrol

import (
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/tyrol/bolzano"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/tyrol/innsbruck"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/tyrol/tirol"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/tyrol/trent"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TyrolDuke interface {
	feud.Duke
	CBolzano博尔扎诺() bolzano.BolzanoCounty
	CInnsbruck因斯布鲁克() innsbruck.InnsbruckCounty
	CTirol兰德克() tirol.TirolCounty
	CTrent特伦托() trent.TrentCounty
}

type 蒂罗尔TyrolDuke struct {
	feud.BaseDuke
	博尔扎诺Bolzano    bolzano.BolzanoCounty
	因斯布鲁克Innsbruck innsbruck.InnsbruckCounty
	兰德克Tirol       tirol.TirolCounty
	特伦托Trent       trent.TrentCounty
}

func (d *蒂罗尔TyrolDuke) CBolzano博尔扎诺() bolzano.BolzanoCounty {
	return d.博尔扎诺Bolzano
}

func (d *蒂罗尔TyrolDuke) CInnsbruck因斯布鲁克() innsbruck.InnsbruckCounty {
	return d.因斯布鲁克Innsbruck
}

func (d *蒂罗尔TyrolDuke) CTirol兰德克() tirol.TirolCounty {
	return d.兰德克Tirol
}

func (d *蒂罗尔TyrolDuke) CTrent特伦托() trent.TrentCounty {
	return d.特伦托Trent
}

var DTyrol蒂罗尔 TyrolDuke = &蒂罗尔TyrolDuke{}

func init() {
	f := DTyrol蒂罗尔.(*蒂罗尔TyrolDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tyrol",
		TitleName: "蒂罗尔",
		TitleCode: "d_tyrol",
		Counties:  map[string]feud.County{},
	}

	f.博尔扎诺Bolzano = bolzano.CBolzano博尔扎诺
	f.博尔扎诺Bolzano.SetParent(f)

	f.因斯布鲁克Innsbruck = innsbruck.CInnsbruck因斯布鲁克
	f.因斯布鲁克Innsbruck.SetParent(f)

	f.兰德克Tirol = tirol.CTirol兰德克
	f.兰德克Tirol.SetParent(f)

	f.特伦托Trent = trent.CTrent特伦托
	f.特伦托Trent.SetParent(f)

}
