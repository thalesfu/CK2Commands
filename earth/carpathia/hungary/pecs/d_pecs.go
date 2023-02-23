package pecs

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/pecs/fejer"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/pecs/pecs"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/pecs/szekezfehervar"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/pecs/vas"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PecsDuke interface {
	feud.Duke
	CFejer费耶尔() fejer.FejerCounty
	CPecs佩奇() pecs.PecsCounty
	CSzekezfehervar塞克什白堡() szekezfehervar.SzekezfehervarCounty
	CVas沃什() vas.VasCounty
}

type 佩奇PecsDuke struct {
	feud.BaseDuke
	费耶尔Fejer            fejer.FejerCounty
	佩奇Pecs              pecs.PecsCounty
	塞克什白堡Szekezfehervar szekezfehervar.SzekezfehervarCounty
	沃什Vas               vas.VasCounty
}

func (d *佩奇PecsDuke) CFejer费耶尔() fejer.FejerCounty {
	return d.费耶尔Fejer
}

func (d *佩奇PecsDuke) CPecs佩奇() pecs.PecsCounty {
	return d.佩奇Pecs
}

func (d *佩奇PecsDuke) CSzekezfehervar塞克什白堡() szekezfehervar.SzekezfehervarCounty {
	return d.塞克什白堡Szekezfehervar
}

func (d *佩奇PecsDuke) CVas沃什() vas.VasCounty {
	return d.沃什Vas
}

var DPecs佩奇 PecsDuke = &佩奇PecsDuke{}

func init() {
	f := DPecs佩奇.(*佩奇PecsDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "pecs",
		TitleName: "佩奇",
		TitleCode: "d_pecs",
		Counties:  map[string]feud.County{},
	}

	f.费耶尔Fejer = fejer.CFejer费耶尔
	f.费耶尔Fejer.SetParent(f)

	f.佩奇Pecs = pecs.CPecs佩奇
	f.佩奇Pecs.SetParent(f)

	f.塞克什白堡Szekezfehervar = szekezfehervar.CSzekezfehervar塞克什白堡
	f.塞克什白堡Szekezfehervar.SetParent(f)

	f.沃什Vas = vas.CVas沃什
	f.沃什Vas.SetParent(f)

}
