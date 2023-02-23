package nasikya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NasikyaCounty interface {
	feud.County
	BIgatpuri伊格德布里() feud.Barony
	BNasikya那私迦() feud.Barony
	BPatta帕塔() feud.Barony
	BRatangad剌怛那揭陀() feud.Barony
	BSaptashrungi七峰() feud.Barony
	BSeunapura犀浮那补罗() feud.Barony
	BTryambaka底哩阎婆迦() feud.Barony
}

type 那私迦NasikyaCounty struct {
	feud.BaseCounty
	伊格德布里Igatpuri  feud.Barony
	那私迦Nasikya     feud.Barony
	帕塔Patta        feud.Barony
	剌怛那揭陀Ratangad  feud.Barony
	七峰Saptashrungi feud.Barony
	犀浮那补罗Seunapura feud.Barony
	底哩阎婆迦Tryambaka feud.Barony
}

func (c *那私迦NasikyaCounty) BIgatpuri伊格德布里() feud.Barony {
	return c.伊格德布里Igatpuri
}

func (c *那私迦NasikyaCounty) BNasikya那私迦() feud.Barony {
	return c.那私迦Nasikya
}

func (c *那私迦NasikyaCounty) BPatta帕塔() feud.Barony {
	return c.帕塔Patta
}

func (c *那私迦NasikyaCounty) BRatangad剌怛那揭陀() feud.Barony {
	return c.剌怛那揭陀Ratangad
}

func (c *那私迦NasikyaCounty) BSaptashrungi七峰() feud.Barony {
	return c.七峰Saptashrungi
}

func (c *那私迦NasikyaCounty) BSeunapura犀浮那补罗() feud.Barony {
	return c.犀浮那补罗Seunapura
}

func (c *那私迦NasikyaCounty) BTryambaka底哩阎婆迦() feud.Barony {
	return c.底哩阎婆迦Tryambaka
}

var CNasikya那私迦 NasikyaCounty = &那私迦NasikyaCounty{}

func init() {
	f := CNasikya那私迦.(*那私迦NasikyaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1260",
		Title:     "nasikya",
		TitleName: "那私迦",
		TitleCode: "c_nasikya",
		Baronies:  map[string]feud.Barony{},
	}

	f.伊格德布里Igatpuri = BIgatpuri伊格德布里
	f.伊格德布里Igatpuri.SetParent(f)

	f.那私迦Nasikya = BNasikya那私迦
	f.那私迦Nasikya.SetParent(f)

	f.帕塔Patta = BPatta帕塔
	f.帕塔Patta.SetParent(f)

	f.剌怛那揭陀Ratangad = BRatangad剌怛那揭陀
	f.剌怛那揭陀Ratangad.SetParent(f)

	f.七峰Saptashrungi = BSaptashrungi七峰
	f.七峰Saptashrungi.SetParent(f)

	f.犀浮那补罗Seunapura = BSeunapura犀浮那补罗
	f.犀浮那补罗Seunapura.SetParent(f)

	f.底哩阎婆迦Tryambaka = BTryambaka底哩阎婆迦
	f.底哩阎婆迦Tryambaka.SetParent(f)

}
