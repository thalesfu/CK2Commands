package basra

import (
	"github.com/thalesfu/CK2Commands/earth/persia/iraq/basra/basra"
	"github.com/thalesfu/CK2Commands/earth/persia/iraq/basra/kufa"
	"github.com/thalesfu/CK2Commands/earth/persia/iraq/basra/rummah"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BasraDuke interface {
	feud.Duke
	CBasra巴士拉() basra.BasraCounty
	CKufa库法() kufa.KufaCounty
	CRummah艾因赛德() rummah.RummahCounty
}

type 巴士拉BasraDuke struct {
	feud.BaseDuke
	巴士拉Basra   basra.BasraCounty
	库法Kufa     kufa.KufaCounty
	艾因赛德Rummah rummah.RummahCounty
}

func (d *巴士拉BasraDuke) CBasra巴士拉() basra.BasraCounty {
	return d.巴士拉Basra
}

func (d *巴士拉BasraDuke) CKufa库法() kufa.KufaCounty {
	return d.库法Kufa
}

func (d *巴士拉BasraDuke) CRummah艾因赛德() rummah.RummahCounty {
	return d.艾因赛德Rummah
}

var DBasra巴士拉 BasraDuke = &巴士拉BasraDuke{}

func init() {
	f := DBasra巴士拉.(*巴士拉BasraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "basra",
		TitleName: "巴士拉",
		TitleCode: "d_basra",
		Counties:  map[string]feud.County{},
	}

	f.巴士拉Basra = basra.CBasra巴士拉
	f.巴士拉Basra.SetParent(f)

	f.库法Kufa = kufa.CKufa库法
	f.库法Kufa.SetParent(f)

	f.艾因赛德Rummah = rummah.CRummah艾因赛德
	f.艾因赛德Rummah.SetParent(f)

}
