package palmyra

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/palmyra/druz"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/palmyra/palmyra"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/palmyra/syria"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/palmyra/tadmor"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PalmyraDuke interface {
	feud.Duke
	CDruz德鲁兹() druz.DruzCounty
	CPalmyra巴尔米拉() palmyra.PalmyraCounty
	CSyria布斯拉() syria.SyriaCounty
	CTadmor苏赫奈() tadmor.TadmorCounty
}

type 巴尔米拉PalmyraDuke struct {
	feud.BaseDuke
	德鲁兹Druz     druz.DruzCounty
	巴尔米拉Palmyra palmyra.PalmyraCounty
	布斯拉Syria    syria.SyriaCounty
	苏赫奈Tadmor   tadmor.TadmorCounty
}

func (d *巴尔米拉PalmyraDuke) CDruz德鲁兹() druz.DruzCounty {
	return d.德鲁兹Druz
}

func (d *巴尔米拉PalmyraDuke) CPalmyra巴尔米拉() palmyra.PalmyraCounty {
	return d.巴尔米拉Palmyra
}

func (d *巴尔米拉PalmyraDuke) CSyria布斯拉() syria.SyriaCounty {
	return d.布斯拉Syria
}

func (d *巴尔米拉PalmyraDuke) CTadmor苏赫奈() tadmor.TadmorCounty {
	return d.苏赫奈Tadmor
}

var DPalmyra巴尔米拉 PalmyraDuke = &巴尔米拉PalmyraDuke{}

func init() {
	f := DPalmyra巴尔米拉.(*巴尔米拉PalmyraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "palmyra",
		TitleName: "巴尔米拉",
		TitleCode: "d_palmyra",
		Counties:  map[string]feud.County{},
	}

	f.德鲁兹Druz = druz.CDruz德鲁兹
	f.德鲁兹Druz.SetParent(f)

	f.巴尔米拉Palmyra = palmyra.CPalmyra巴尔米拉
	f.巴尔米拉Palmyra.SetParent(f)

	f.布斯拉Syria = syria.CSyria布斯拉
	f.布斯拉Syria.SetParent(f)

	f.苏赫奈Tadmor = tadmor.CTadmor苏赫奈
	f.苏赫奈Tadmor.SetParent(f)

}
