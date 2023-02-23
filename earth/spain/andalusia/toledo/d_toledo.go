package toledo

import (
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/toledo/alcala"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/toledo/cuenca"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/toledo/molina"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/toledo/toledo"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ToledoDuke interface {
	feud.Duke
	CAlcala阿尔卡拉() alcala.AlcalaCounty
	CCuenca昆卡() cuenca.CuencaCounty
	CMolina莫利纳() molina.MolinaCounty
	CToledo托雷多() toledo.ToledoCounty
}

type 托雷多ToledoDuke struct {
	feud.BaseDuke
	阿尔卡拉Alcala alcala.AlcalaCounty
	昆卡Cuenca   cuenca.CuencaCounty
	莫利纳Molina  molina.MolinaCounty
	托雷多Toledo  toledo.ToledoCounty
}

func (d *托雷多ToledoDuke) CAlcala阿尔卡拉() alcala.AlcalaCounty {
	return d.阿尔卡拉Alcala
}

func (d *托雷多ToledoDuke) CCuenca昆卡() cuenca.CuencaCounty {
	return d.昆卡Cuenca
}

func (d *托雷多ToledoDuke) CMolina莫利纳() molina.MolinaCounty {
	return d.莫利纳Molina
}

func (d *托雷多ToledoDuke) CToledo托雷多() toledo.ToledoCounty {
	return d.托雷多Toledo
}

var DToledo托雷多 ToledoDuke = &托雷多ToledoDuke{}

func init() {
	f := DToledo托雷多.(*托雷多ToledoDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "toledo",
		TitleName: "托雷多",
		TitleCode: "d_toledo",
		Counties:  map[string]feud.County{},
	}

	f.阿尔卡拉Alcala = alcala.CAlcala阿尔卡拉
	f.阿尔卡拉Alcala.SetParent(f)

	f.昆卡Cuenca = cuenca.CCuenca昆卡
	f.昆卡Cuenca.SetParent(f)

	f.莫利纳Molina = molina.CMolina莫利纳
	f.莫利纳Molina.SetParent(f)

	f.托雷多Toledo = toledo.CToledo托雷多
	f.托雷多Toledo.SetParent(f)

}
