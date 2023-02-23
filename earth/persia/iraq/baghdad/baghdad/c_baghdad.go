package baghdad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BaghdadCounty interface {
	feud.County
	BBabel巴贝尔() feud.Barony
	BBagdad巴格达() feud.Barony
	BBaqubah巴古拜() feud.Barony
	BHillah希拉() feud.Barony
	BIskandariya伊斯坎德里耶() feud.Barony
	BLatifiya拉蒂菲亚() feud.Barony
	BMadain马达因() feud.Barony
	BTaji塔吉() feud.Barony
}

type 巴格达BaghdadCounty struct {
	feud.BaseCounty
	巴贝尔Babel          feud.Barony
	巴格达Bagdad         feud.Barony
	巴古拜Baqubah        feud.Barony
	希拉Hillah          feud.Barony
	伊斯坎德里耶Iskandariya feud.Barony
	拉蒂菲亚Latifiya      feud.Barony
	马达因Madain         feud.Barony
	塔吉Taji            feud.Barony
}

func (c *巴格达BaghdadCounty) BBabel巴贝尔() feud.Barony {
	return c.巴贝尔Babel
}

func (c *巴格达BaghdadCounty) BBagdad巴格达() feud.Barony {
	return c.巴格达Bagdad
}

func (c *巴格达BaghdadCounty) BBaqubah巴古拜() feud.Barony {
	return c.巴古拜Baqubah
}

func (c *巴格达BaghdadCounty) BHillah希拉() feud.Barony {
	return c.希拉Hillah
}

func (c *巴格达BaghdadCounty) BIskandariya伊斯坎德里耶() feud.Barony {
	return c.伊斯坎德里耶Iskandariya
}

func (c *巴格达BaghdadCounty) BLatifiya拉蒂菲亚() feud.Barony {
	return c.拉蒂菲亚Latifiya
}

func (c *巴格达BaghdadCounty) BMadain马达因() feud.Barony {
	return c.马达因Madain
}

func (c *巴格达BaghdadCounty) BTaji塔吉() feud.Barony {
	return c.塔吉Taji
}

var CBaghdad巴格达 BaghdadCounty = &巴格达BaghdadCounty{}

func init() {
	f := CBaghdad巴格达.(*巴格达BaghdadCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "693",
		Title:     "baghdad",
		TitleName: "巴格达",
		TitleCode: "c_baghdad",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴贝尔Babel = BBabel巴贝尔
	f.巴贝尔Babel.SetParent(f)

	f.巴格达Bagdad = BBagdad巴格达
	f.巴格达Bagdad.SetParent(f)

	f.巴古拜Baqubah = BBaqubah巴古拜
	f.巴古拜Baqubah.SetParent(f)

	f.希拉Hillah = BHillah希拉
	f.希拉Hillah.SetParent(f)

	f.伊斯坎德里耶Iskandariya = BIskandariya伊斯坎德里耶
	f.伊斯坎德里耶Iskandariya.SetParent(f)

	f.拉蒂菲亚Latifiya = BLatifiya拉蒂菲亚
	f.拉蒂菲亚Latifiya.SetParent(f)

	f.马达因Madain = BMadain马达因
	f.马达因Madain.SetParent(f)

	f.塔吉Taji = BTaji塔吉
	f.塔吉Taji.SetParent(f)

}
