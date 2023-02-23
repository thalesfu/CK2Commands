package jharkand

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/jharkand/jharkand"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/jharkand/munda"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/jharkand/rajrappa"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/jharkand/rothas"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JharkandDuke interface {
	feud.Duke
	CJharkand阇罗犍度() jharkand.JharkandCounty
	CMunda文荼() munda.MundaCounty
	CRajrappa罗奢罗波() rajrappa.RajrappaCounty
	CRothas卢诃多娑() rothas.RothasCounty
}

type 阇罗犍度JharkandDuke struct {
	feud.BaseDuke
	阇罗犍度Jharkand jharkand.JharkandCounty
	文荼Munda      munda.MundaCounty
	罗奢罗波Rajrappa rajrappa.RajrappaCounty
	卢诃多娑Rothas   rothas.RothasCounty
}

func (d *阇罗犍度JharkandDuke) CJharkand阇罗犍度() jharkand.JharkandCounty {
	return d.阇罗犍度Jharkand
}

func (d *阇罗犍度JharkandDuke) CMunda文荼() munda.MundaCounty {
	return d.文荼Munda
}

func (d *阇罗犍度JharkandDuke) CRajrappa罗奢罗波() rajrappa.RajrappaCounty {
	return d.罗奢罗波Rajrappa
}

func (d *阇罗犍度JharkandDuke) CRothas卢诃多娑() rothas.RothasCounty {
	return d.卢诃多娑Rothas
}

var DJharkand阇罗犍度 JharkandDuke = &阇罗犍度JharkandDuke{}

func init() {
	f := DJharkand阇罗犍度.(*阇罗犍度JharkandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "jharkand",
		TitleName: "阇罗犍度",
		TitleCode: "d_jharkand",
		Counties:  map[string]feud.County{},
	}

	f.阇罗犍度Jharkand = jharkand.CJharkand阇罗犍度
	f.阇罗犍度Jharkand.SetParent(f)

	f.文荼Munda = munda.CMunda文荼
	f.文荼Munda.SetParent(f)

	f.罗奢罗波Rajrappa = rajrappa.CRajrappa罗奢罗波
	f.罗奢罗波Rajrappa.SetParent(f)

	f.卢诃多娑Rothas = rothas.CRothas卢诃多娑
	f.卢诃多娑Rothas.SetParent(f)

}
