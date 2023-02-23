package polotsk

import (
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/polotsk/lepiel"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/polotsk/orsha"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/polotsk/polotsk"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/polotsk/vitebsk"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PolotskDuke interface {
	feud.Duke
	CLepiel卢科姆利() lepiel.LepielCounty
	COrsha奥尔沙() orsha.OrshaCounty
	CPolotsk波洛茨克() polotsk.PolotskCounty
	CVitebsk维捷布斯克() vitebsk.VitebskCounty
}

type 波洛茨克PolotskDuke struct {
	feud.BaseDuke
	卢科姆利Lepiel   lepiel.LepielCounty
	奥尔沙Orsha     orsha.OrshaCounty
	波洛茨克Polotsk  polotsk.PolotskCounty
	维捷布斯克Vitebsk vitebsk.VitebskCounty
}

func (d *波洛茨克PolotskDuke) CLepiel卢科姆利() lepiel.LepielCounty {
	return d.卢科姆利Lepiel
}

func (d *波洛茨克PolotskDuke) COrsha奥尔沙() orsha.OrshaCounty {
	return d.奥尔沙Orsha
}

func (d *波洛茨克PolotskDuke) CPolotsk波洛茨克() polotsk.PolotskCounty {
	return d.波洛茨克Polotsk
}

func (d *波洛茨克PolotskDuke) CVitebsk维捷布斯克() vitebsk.VitebskCounty {
	return d.维捷布斯克Vitebsk
}

var DPolotsk波洛茨克 PolotskDuke = &波洛茨克PolotskDuke{}

func init() {
	f := DPolotsk波洛茨克.(*波洛茨克PolotskDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "polotsk",
		TitleName: "波洛茨克",
		TitleCode: "d_polotsk",
		Counties:  map[string]feud.County{},
	}

	f.卢科姆利Lepiel = lepiel.CLepiel卢科姆利
	f.卢科姆利Lepiel.SetParent(f)

	f.奥尔沙Orsha = orsha.COrsha奥尔沙
	f.奥尔沙Orsha.SetParent(f)

	f.波洛茨克Polotsk = polotsk.CPolotsk波洛茨克
	f.波洛茨克Polotsk.SetParent(f)

	f.维捷布斯克Vitebsk = vitebsk.CVitebsk维捷布斯克
	f.维捷布斯克Vitebsk.SetParent(f)

}
