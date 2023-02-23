package pisa

import (
	"github.com/thalesfu/CK2Commands/earth/italy/italy/pisa/piombino"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/pisa/pisa"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PisaDuke interface {
	feud.Duke
	CPiombino皮翁比诺() piombino.PiombinoCounty
	CPisa比萨() pisa.PisaCounty
}

type 比萨PisaDuke struct {
	feud.BaseDuke
	皮翁比诺Piombino piombino.PiombinoCounty
	比萨Pisa       pisa.PisaCounty
}

func (d *比萨PisaDuke) CPiombino皮翁比诺() piombino.PiombinoCounty {
	return d.皮翁比诺Piombino
}

func (d *比萨PisaDuke) CPisa比萨() pisa.PisaCounty {
	return d.比萨Pisa
}

var DPisa比萨 PisaDuke = &比萨PisaDuke{}

func init() {
	f := DPisa比萨.(*比萨PisaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "pisa",
		TitleName: "比萨",
		TitleCode: "d_pisa",
		Counties:  map[string]feud.County{},
	}

	f.皮翁比诺Piombino = piombino.CPiombino皮翁比诺
	f.皮翁比诺Piombino.SetParent(f)

	f.比萨Pisa = pisa.CPisa比萨
	f.比萨Pisa.SetParent(f)

}
