package shiraz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ShirazCounty interface {
	feud.County
	BAbadeh阿巴代() feud.Barony
	BArsanjan阿桑詹() feud.Barony
	BAzargarta埃格利德() feud.Barony
	BBavanat巴瓦纳特() feud.Barony
	BBishapur比沙普尔() feud.Barony
	BEstakhr耶斯塔赫尔() feud.Barony
	BPersepolis波斯波利斯() feud.Barony
	BShiraz设拉子() feud.Barony
}

type 设拉子ShirazCounty struct {
	feud.BaseCounty
	阿巴代Abadeh       feud.Barony
	阿桑詹Arsanjan     feud.Barony
	埃格利德Azargarta   feud.Barony
	巴瓦纳特Bavanat     feud.Barony
	比沙普尔Bishapur    feud.Barony
	耶斯塔赫尔Estakhr    feud.Barony
	波斯波利斯Persepolis feud.Barony
	设拉子Shiraz       feud.Barony
}

func (c *设拉子ShirazCounty) BAbadeh阿巴代() feud.Barony {
	return c.阿巴代Abadeh
}

func (c *设拉子ShirazCounty) BArsanjan阿桑詹() feud.Barony {
	return c.阿桑詹Arsanjan
}

func (c *设拉子ShirazCounty) BAzargarta埃格利德() feud.Barony {
	return c.埃格利德Azargarta
}

func (c *设拉子ShirazCounty) BBavanat巴瓦纳特() feud.Barony {
	return c.巴瓦纳特Bavanat
}

func (c *设拉子ShirazCounty) BBishapur比沙普尔() feud.Barony {
	return c.比沙普尔Bishapur
}

func (c *设拉子ShirazCounty) BEstakhr耶斯塔赫尔() feud.Barony {
	return c.耶斯塔赫尔Estakhr
}

func (c *设拉子ShirazCounty) BPersepolis波斯波利斯() feud.Barony {
	return c.波斯波利斯Persepolis
}

func (c *设拉子ShirazCounty) BShiraz设拉子() feud.Barony {
	return c.设拉子Shiraz
}

var CShiraz设拉子 ShirazCounty = &设拉子ShirazCounty{}

func init() {
	f := CShiraz设拉子.(*设拉子ShirazCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "644",
		Title:     "shiraz",
		TitleName: "设拉子",
		TitleCode: "c_shiraz",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴代Abadeh = BAbadeh阿巴代
	f.阿巴代Abadeh.SetParent(f)

	f.阿桑詹Arsanjan = BArsanjan阿桑詹
	f.阿桑詹Arsanjan.SetParent(f)

	f.埃格利德Azargarta = BAzargarta埃格利德
	f.埃格利德Azargarta.SetParent(f)

	f.巴瓦纳特Bavanat = BBavanat巴瓦纳特
	f.巴瓦纳特Bavanat.SetParent(f)

	f.比沙普尔Bishapur = BBishapur比沙普尔
	f.比沙普尔Bishapur.SetParent(f)

	f.耶斯塔赫尔Estakhr = BEstakhr耶斯塔赫尔
	f.耶斯塔赫尔Estakhr.SetParent(f)

	f.波斯波利斯Persepolis = BPersepolis波斯波利斯
	f.波斯波利斯Persepolis.SetParent(f)

	f.设拉子Shiraz = BShiraz设拉子
	f.设拉子Shiraz.SetParent(f)

}
