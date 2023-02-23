package gurma

import (
	"github.com/thalesfu/CK2Commands/earth/mali/songhay/gurma/dendi"
	"github.com/thalesfu/CK2Commands/earth/mali/songhay/gurma/gurma"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GurmaDuke interface {
	feud.Duke
	CDendi登迪() dendi.DendiCounty
	CGurma古尔玛() gurma.GurmaCounty
}

type 古尔马GurmaDuke struct {
	feud.BaseDuke
	登迪Dendi  dendi.DendiCounty
	古尔玛Gurma gurma.GurmaCounty
}

func (d *古尔马GurmaDuke) CDendi登迪() dendi.DendiCounty {
	return d.登迪Dendi
}

func (d *古尔马GurmaDuke) CGurma古尔玛() gurma.GurmaCounty {
	return d.古尔玛Gurma
}

var DGurma古尔马 GurmaDuke = &古尔马GurmaDuke{}

func init() {
	f := DGurma古尔马.(*古尔马GurmaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gurma",
		TitleName: "古尔马",
		TitleCode: "d_gurma",
		Counties:  map[string]feud.County{},
	}

	f.登迪Dendi = dendi.CDendi登迪
	f.登迪Dendi.SetParent(f)

	f.古尔玛Gurma = gurma.CGurma古尔玛
	f.古尔玛Gurma.SetParent(f)

}
