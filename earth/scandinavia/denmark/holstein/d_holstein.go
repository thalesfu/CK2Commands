package holstein

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/holstein/dithmarschen"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/holstein/holstein"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/holstein/slesvig"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HolsteinDuke interface {
	feud.Duke
	CDithmarschen迪特马申() dithmarschen.DithmarschenCounty
	CHolstein荷尔斯泰因() holstein.HolsteinCounty
	CSlesvig石勒苏益格() slesvig.SlesvigCounty
}

type 荷尔斯泰因HolsteinDuke struct {
	feud.BaseDuke
	迪特马申Dithmarschen dithmarschen.DithmarschenCounty
	荷尔斯泰因Holstein    holstein.HolsteinCounty
	石勒苏益格Slesvig     slesvig.SlesvigCounty
}

func (d *荷尔斯泰因HolsteinDuke) CDithmarschen迪特马申() dithmarschen.DithmarschenCounty {
	return d.迪特马申Dithmarschen
}

func (d *荷尔斯泰因HolsteinDuke) CHolstein荷尔斯泰因() holstein.HolsteinCounty {
	return d.荷尔斯泰因Holstein
}

func (d *荷尔斯泰因HolsteinDuke) CSlesvig石勒苏益格() slesvig.SlesvigCounty {
	return d.石勒苏益格Slesvig
}

var DHolstein荷尔斯泰因 HolsteinDuke = &荷尔斯泰因HolsteinDuke{}

func init() {
	f := DHolstein荷尔斯泰因.(*荷尔斯泰因HolsteinDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "holstein",
		TitleName: "荷尔斯泰因",
		TitleCode: "d_holstein",
		Counties:  map[string]feud.County{},
	}

	f.迪特马申Dithmarschen = dithmarschen.CDithmarschen迪特马申
	f.迪特马申Dithmarschen.SetParent(f)

	f.荷尔斯泰因Holstein = holstein.CHolstein荷尔斯泰因
	f.荷尔斯泰因Holstein.SetParent(f)

	f.石勒苏益格Slesvig = slesvig.CSlesvig石勒苏益格
	f.石勒苏益格Slesvig.SetParent(f)

}
