package tver

import (
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/tver/kashin"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/tver/mozhaysk"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/tver/rzheva"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/tver/tver"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/tver/zoubtsov"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TverDuke interface {
	feud.Duke
	CKashin卡申() kashin.KashinCounty
	CMozhaysk莫扎伊斯克() mozhaysk.MozhayskCounty
	CRzheva勒热瓦() rzheva.RzhevaCounty
	CTver特维尔() tver.TverCounty
	CZoubtsov祖布佐夫() zoubtsov.ZoubtsovCounty
}

type 特维尔TverDuke struct {
	feud.BaseDuke
	卡申Kashin      kashin.KashinCounty
	莫扎伊斯克Mozhaysk mozhaysk.MozhayskCounty
	勒热瓦Rzheva     rzheva.RzhevaCounty
	特维尔Tver       tver.TverCounty
	祖布佐夫Zoubtsov  zoubtsov.ZoubtsovCounty
}

func (d *特维尔TverDuke) CKashin卡申() kashin.KashinCounty {
	return d.卡申Kashin
}

func (d *特维尔TverDuke) CMozhaysk莫扎伊斯克() mozhaysk.MozhayskCounty {
	return d.莫扎伊斯克Mozhaysk
}

func (d *特维尔TverDuke) CRzheva勒热瓦() rzheva.RzhevaCounty {
	return d.勒热瓦Rzheva
}

func (d *特维尔TverDuke) CTver特维尔() tver.TverCounty {
	return d.特维尔Tver
}

func (d *特维尔TverDuke) CZoubtsov祖布佐夫() zoubtsov.ZoubtsovCounty {
	return d.祖布佐夫Zoubtsov
}

var DTver特维尔 TverDuke = &特维尔TverDuke{}

func init() {
	f := DTver特维尔.(*特维尔TverDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tver",
		TitleName: "特维尔",
		TitleCode: "d_tver",
		Counties:  map[string]feud.County{},
	}

	f.卡申Kashin = kashin.CKashin卡申
	f.卡申Kashin.SetParent(f)

	f.莫扎伊斯克Mozhaysk = mozhaysk.CMozhaysk莫扎伊斯克
	f.莫扎伊斯克Mozhaysk.SetParent(f)

	f.勒热瓦Rzheva = rzheva.CRzheva勒热瓦
	f.勒热瓦Rzheva.SetParent(f)

	f.特维尔Tver = tver.CTver特维尔
	f.特维尔Tver.SetParent(f)

	f.祖布佐夫Zoubtsov = zoubtsov.CZoubtsov祖布佐夫
	f.祖布佐夫Zoubtsov.SetParent(f)

}
