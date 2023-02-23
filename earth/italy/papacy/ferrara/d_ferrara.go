package ferrara

import (
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/ferrara/bologna"
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/ferrara/ferrara"
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/ferrara/ravenna"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FerraraDuke interface {
	feud.Duke
	CBologna博洛尼亚() bologna.BolognaCounty
	CFerrara费拉拉() ferrara.FerraraCounty
	CRavenna拉文纳() ravenna.RavennaCounty
}

type 费拉拉FerraraDuke struct {
	feud.BaseDuke
	博洛尼亚Bologna bologna.BolognaCounty
	费拉拉Ferrara  ferrara.FerraraCounty
	拉文纳Ravenna  ravenna.RavennaCounty
}

func (d *费拉拉FerraraDuke) CBologna博洛尼亚() bologna.BolognaCounty {
	return d.博洛尼亚Bologna
}

func (d *费拉拉FerraraDuke) CFerrara费拉拉() ferrara.FerraraCounty {
	return d.费拉拉Ferrara
}

func (d *费拉拉FerraraDuke) CRavenna拉文纳() ravenna.RavennaCounty {
	return d.拉文纳Ravenna
}

var DFerrara费拉拉 FerraraDuke = &费拉拉FerraraDuke{}

func init() {
	f := DFerrara费拉拉.(*费拉拉FerraraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ferrara",
		TitleName: "费拉拉",
		TitleCode: "d_ferrara",
		Counties:  map[string]feud.County{},
	}

	f.博洛尼亚Bologna = bologna.CBologna博洛尼亚
	f.博洛尼亚Bologna.SetParent(f)

	f.费拉拉Ferrara = ferrara.CFerrara费拉拉
	f.费拉拉Ferrara.SetParent(f)

	f.拉文纳Ravenna = ravenna.CRavenna拉文纳
	f.拉文纳Ravenna.SetParent(f)

}
