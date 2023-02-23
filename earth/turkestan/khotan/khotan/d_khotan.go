package khotan

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/khotan/cadota"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/khotan/cherchen"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/khotan/karghalik"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/khotan/keriya"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/khotan/khotan"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/khotan/yarkand"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhotanDuke interface {
	feud.Duke
	CCadota精绝() cadota.CadotaCounty
	CCherchen车尔臣() cherchen.CherchenCounty
	CKarghalik喀格勒克() karghalik.KarghalikCounty
	CKeriya克里雅() keriya.KeriyaCounty
	CKhotan于阗() khotan.KhotanCounty
	CYarkand鸦儿看() yarkand.YarkandCounty
}

type 于阗KhotanDuke struct {
	feud.BaseDuke
	精绝Cadota      cadota.CadotaCounty
	车尔臣Cherchen   cherchen.CherchenCounty
	喀格勒克Karghalik karghalik.KarghalikCounty
	克里雅Keriya     keriya.KeriyaCounty
	于阗Khotan      khotan.KhotanCounty
	鸦儿看Yarkand    yarkand.YarkandCounty
}

func (d *于阗KhotanDuke) CCadota精绝() cadota.CadotaCounty {
	return d.精绝Cadota
}

func (d *于阗KhotanDuke) CCherchen车尔臣() cherchen.CherchenCounty {
	return d.车尔臣Cherchen
}

func (d *于阗KhotanDuke) CKarghalik喀格勒克() karghalik.KarghalikCounty {
	return d.喀格勒克Karghalik
}

func (d *于阗KhotanDuke) CKeriya克里雅() keriya.KeriyaCounty {
	return d.克里雅Keriya
}

func (d *于阗KhotanDuke) CKhotan于阗() khotan.KhotanCounty {
	return d.于阗Khotan
}

func (d *于阗KhotanDuke) CYarkand鸦儿看() yarkand.YarkandCounty {
	return d.鸦儿看Yarkand
}

var DKhotan于阗 KhotanDuke = &于阗KhotanDuke{}

func init() {
	f := DKhotan于阗.(*于阗KhotanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "khotan",
		TitleName: "于阗",
		TitleCode: "d_khotan",
		Counties:  map[string]feud.County{},
	}

	f.精绝Cadota = cadota.CCadota精绝
	f.精绝Cadota.SetParent(f)

	f.车尔臣Cherchen = cherchen.CCherchen车尔臣
	f.车尔臣Cherchen.SetParent(f)

	f.喀格勒克Karghalik = karghalik.CKarghalik喀格勒克
	f.喀格勒克Karghalik.SetParent(f)

	f.克里雅Keriya = keriya.CKeriya克里雅
	f.克里雅Keriya.SetParent(f)

	f.于阗Khotan = khotan.CKhotan于阗
	f.于阗Khotan.SetParent(f)

	f.鸦儿看Yarkand = yarkand.CYarkand鸦儿看
	f.鸦儿看Yarkand.SetParent(f)

}
