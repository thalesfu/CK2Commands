package turov

import (
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/turov/pinsk"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/turov/slutsk"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/turov/turov"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/turov/volkovysk"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TurovDuke interface {
	feud.Duke
	CPinsk平斯克() pinsk.PinskCounty
	CSlutsk斯卢茨克() slutsk.SlutskCounty
	CTurov图罗夫() turov.TurovCounty
	CVolkovysk沃尔科维斯克() volkovysk.VolkovyskCounty
}

type 图罗夫TurovDuke struct {
	feud.BaseDuke
	平斯克Pinsk        pinsk.PinskCounty
	斯卢茨克Slutsk      slutsk.SlutskCounty
	图罗夫Turov        turov.TurovCounty
	沃尔科维斯克Volkovysk volkovysk.VolkovyskCounty
}

func (d *图罗夫TurovDuke) CPinsk平斯克() pinsk.PinskCounty {
	return d.平斯克Pinsk
}

func (d *图罗夫TurovDuke) CSlutsk斯卢茨克() slutsk.SlutskCounty {
	return d.斯卢茨克Slutsk
}

func (d *图罗夫TurovDuke) CTurov图罗夫() turov.TurovCounty {
	return d.图罗夫Turov
}

func (d *图罗夫TurovDuke) CVolkovysk沃尔科维斯克() volkovysk.VolkovyskCounty {
	return d.沃尔科维斯克Volkovysk
}

var DTurov图罗夫 TurovDuke = &图罗夫TurovDuke{}

func init() {
	f := DTurov图罗夫.(*图罗夫TurovDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "turov",
		TitleName: "图罗夫",
		TitleCode: "d_turov",
		Counties:  map[string]feud.County{},
	}

	f.平斯克Pinsk = pinsk.CPinsk平斯克
	f.平斯克Pinsk.SetParent(f)

	f.斯卢茨克Slutsk = slutsk.CSlutsk斯卢茨克
	f.斯卢茨克Slutsk.SetParent(f)

	f.图罗夫Turov = turov.CTurov图罗夫
	f.图罗夫Turov.SetParent(f)

	f.沃尔科维斯克Volkovysk = volkovysk.CVolkovysk沃尔科维斯克
	f.沃尔科维斯克Volkovysk.SetParent(f)

}
