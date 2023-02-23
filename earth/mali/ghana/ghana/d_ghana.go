package ghana

import (
	"github.com/thalesfu/CK2Commands/earth/mali/ghana/ghana/djenne"
	"github.com/thalesfu/CK2Commands/earth/mali/ghana/ghana/ghana"
	"github.com/thalesfu/CK2Commands/earth/mali/ghana/ghana/mema"
	"github.com/thalesfu/CK2Commands/earth/mali/ghana/ghana/soso"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GhanaDuke interface {
	feud.Duke
	CDjenne杰内() djenne.DjenneCounty
	CGhana瓦加杜() ghana.GhanaCounty
	CMema梅马() mema.MemaCounty
	CSoso索索() soso.SosoCounty
}

type 加纳GhanaDuke struct {
	feud.BaseDuke
	杰内Djenne djenne.DjenneCounty
	瓦加杜Ghana ghana.GhanaCounty
	梅马Mema   mema.MemaCounty
	索索Soso   soso.SosoCounty
}

func (d *加纳GhanaDuke) CDjenne杰内() djenne.DjenneCounty {
	return d.杰内Djenne
}

func (d *加纳GhanaDuke) CGhana瓦加杜() ghana.GhanaCounty {
	return d.瓦加杜Ghana
}

func (d *加纳GhanaDuke) CMema梅马() mema.MemaCounty {
	return d.梅马Mema
}

func (d *加纳GhanaDuke) CSoso索索() soso.SosoCounty {
	return d.索索Soso
}

var DGhana加纳 GhanaDuke = &加纳GhanaDuke{}

func init() {
	f := DGhana加纳.(*加纳GhanaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ghana",
		TitleName: "加纳",
		TitleCode: "d_ghana",
		Counties:  map[string]feud.County{},
	}

	f.杰内Djenne = djenne.CDjenne杰内
	f.杰内Djenne.SetParent(f)

	f.瓦加杜Ghana = ghana.CGhana瓦加杜
	f.瓦加杜Ghana.SetParent(f)

	f.梅马Mema = mema.CMema梅马
	f.梅马Mema.SetParent(f)

	f.索索Soso = soso.CSoso索索
	f.索索Soso.SetParent(f)

}
